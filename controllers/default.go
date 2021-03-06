package controllers

import (
	"saml_idhub_api/go-saml"
	"saml_idhub_api/keypairs"
	"saml_idhub_api/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"bufio"
	// "encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/goware/saml"
)

var privateKeyDirPrefix = "./privateKey/"
var certificateDirPrefix = "./certificate/"

type NewError struct {
	Err string
}

type Res struct {
	Bool     bool
	Response string
}

type OrganizationController struct {
	beego.Controller
}

type MetadataController struct {
	beego.Controller
}

type SamlResponseController struct {
	beego.Controller
}

func readCertificate(dir string) (string, error) {
	fi, err := os.Open(dir)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return "", err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	value := ""
	for {
		a, _, c := br.ReadLine()
		if string(a) == "-----BEGIN CERTIFICATE-----" {
			continue
		}
		if string(a) == "-----END CERTIFICATE-----" {
			break
		}
		if c == io.EOF {
			break
		}
		value = value + string(a)
	}
	return value, nil
}

func (c *OrganizationController) Get() {
	id := c.Ctx.Input.Param(":id")
	privateKeyDir := privateKeyDirPrefix + id + ".pem"
	certificateDir := certificateDirPrefix + id + ".crt"
	s, err := keypairs.GenKeyAndCert(privateKeyDir, certificateDir)
	if err != nil {
		fmt.Println(err)
		res := &NewError{err.Error()}
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	fmt.Println(s)

	privateKey, err := ioutil.ReadFile(privateKeyDir)
	if err != nil {
		fmt.Println(err)
		res := &NewError{err.Error()}
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	// fmt.Println(privateKey)

	certificate, err := ioutil.ReadFile(certificateDir)
	if err != nil {
		fmt.Println(err)
		res := &NewError{err.Error()}
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	// fmt.Println(certificate)

	var organization models.Organization
	organization.Key = id
	organization.PrivateKey = string(privateKey)
	organization.Certificate = string(certificate)

	o := orm.NewOrm()
	i, err := o.Insert(&organization)
	fmt.Println(i)
	if err != nil {
		fmt.Println(err)
		res := &NewError{err.Error()}
		c.Data["json"] = res
		c.ServeJSON()
		return
		// 查询数据库失败
	}
	res := &Res{true, "Generate organization success."}
	fmt.Println(res)
	c.Data["json"] = res
	c.ServeJSON()
	return
}

func (c *MetadataController) Get() {
	id := c.Ctx.Input.Param(":id")
	privateKeyDir := privateKeyDirPrefix + id + ".pem"
	certificateDir := certificateDirPrefix + id + ".crt"
	identityProvider := saml.IdentityProvider{
		CertFile: certificateDir,
		KeyFile:  privateKeyDir,

		MetadataURL: "https://saml.idhub.network",
		SSOURL:      "https://samlSSO.idhub.network",

		// SPMetadataURL: "https://saml.idhub.network",
		EntityID: id,

		SecurityOpts: saml.SecurityOpts{
			AllowSelfSignedCert: true,
		},
	}
	// fmt.Println(identityProvider)

	metadata, err := identityProvider.Metadata()
	if err != nil {
		fmt.Println(err)
		res := &NewError{err.Error()}
		c.Data["json"] = res
		c.ServeJSON()
		// Logf("Failed to generate metadata: %v", err)
		// writeErr(w, err)
		return
	}
	c.Data["xml"] = metadata
	fmt.Println(metadata)
	c.ServeXML()
	return
	// out, err := xml.MarshalIndent(metadata, "", "\t")
}

func (c *SamlResponseController) Post() {
	privateKeyDir := privateKeyDirPrefix + c.GetString("org_id") + ".pem"
	certificateDir := certificateDirPrefix + c.GetString("org_id") + ".crt"
	certificate, err := readCertificate(certificateDir)
	if err != nil {
		fmt.Println(err)
		res := &NewError{err.Error()}
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	stringValueOfCert := string(certificate) // "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"
	// fmt.Println(stringValueOfCert)
	issuer := c.GetString("issuer") // "urn:zaakin.idhub.com"
	// fmt.Println(issuer)
	authnRequestIdRespondingTo := "https://signin.aws.amazon.com/saml"
	authnResponse := gosaml.NewSignedResponse()
	authnResponse.Issuer.Url = issuer
	authnResponse.Assertion.Issuer.Url = issuer
	authnResponse.Assertion.Signature.KeyInfo.X509Data.X509Certificate.Cert = stringValueOfCert
	// fmt.Println(user)
	authnResponse.Assertion.Subject.NameID.Value = c.GetString("user_id")
	// fmt.Println(authnResponse.Assertion.Subject.NameID.Value)

	// authnResponse.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.InResponseTo = authnRequestIdRespondingTo
	// authnResponse.InResponseTo = authnRequestIdRespondingTo
	authnResponse.Destination = authnRequestIdRespondingTo
	authnResponse.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.Recipient = authnRequestIdRespondingTo

	authnResponse.AddAudienceRestriction(authnRequestIdRespondingTo)
	authnResponse.AddAttribute("https://aws.amazon.com/SAML/Attributes/Role", c.GetString("role_arn")+","+c.GetString("provider_arn"))
	authnResponse.AddAttribute("https://aws.amazon.com/SAML/Attributes/RoleSessionName", c.GetString("org_aws_id"))
	authnResponse.AddAuthnStatement("urn:oasis:names:tc:SAML:2.0:ac:classes:unspecified", c.GetString("user_id"))

	// fmt.Println(authnResponse)

	// fmt.Println("unsigned!!")

	b64XML, err := authnResponse.EncodedSignedString(privateKeyDir)
	if err != nil {
		fmt.Println(err)
		res := &NewError{err.Error()}
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	// fmt.Println(b64XML, err)
	res := &Res{true, b64XML}
	fmt.Println(res)
	c.Data["json"] = res
	c.ServeJSON()
	return
}

/*
func (c *SamlResponseController) Get() {
	user := "piouslove"
	stringValueOfCert := "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"
	issuer := "urn:zaakin.idhub.com"
	authnRequestIdRespondingTo := "https://signin.aws.amazon.com/saml"
	authnResponse := gosaml.NewSignedResponse()
	authnResponse.Issuer.Url = issuer
	authnResponse.Assertion.Issuer.Url = issuer
	authnResponse.Assertion.Signature.KeyInfo.X509Data.X509Certificate.Cert = stringValueOfCert
	// fmt.Println(user)
	authnResponse.Assertion.Subject.NameID.Value = user
	fmt.Println(authnResponse.Assertion.Subject.NameID.Value)

	// authnResponse.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.InResponseTo = authnRequestIdRespondingTo
	// authnResponse.InResponseTo = authnRequestIdRespondingTo
	authnResponse.Destination = authnRequestIdRespondingTo
	authnResponse.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.Recipient = authnRequestIdRespondingTo

	authnResponse.AddAudienceRestriction(authnRequestIdRespondingTo)
	authnResponse.AddAttribute("https://aws.amazon.com/SAML/Attributes/Role", "arn:aws:iam::434228681164:role/test,arn:aws:iam::434228681164:saml-provider/test")
	authnResponse.AddAttribute("https://aws.amazon.com/SAML/Attributes/RoleSessionName", "434228681164")
	authnResponse.AddAuthnStatement("urn:oasis:names:tc:SAML:2.0:ac:classes:unspecified", "_pF7nFr2Po5NIhvShbHmJAJCjghZa1J8v")
	// fmt.Println("unsigned!!")

	// signed XML string
	signed, err := authnResponse.SignedString("./privateKey/zaakin.pem")
	fmt.Println(signed, err)

	// or signed base64 encoded XML string
	b64XML, err := authnResponse.EncodedSignedString("./privateKey/zaakin.pem")
	// fmt.Println(b64XML, err)
	if err != nil {
		fmt.Println("error:", err)
		res := &NewError{err.Error()}
		c.Data["json"] = res
		c.ServeJSON()
		return
	}
	res := &Res{true, b64XML}
	c.Data["json"] = res
	c.ServeJSON()
	return
}
*/
