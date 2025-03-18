package payslip

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

type payslipConfig struct {
	image string
}

/** NOTE
* One pay group has one setting and meny employees with differnet details and components
* varibale cmptCnfg can be shared across all payslips and need not be cloned
*
* Acctual challanges
* 1. Filtering Employee Data
* Employee Details will probably be in array of fields formate,
* filtering data from it might be expensive as each can be of around 4-5 mb each
* for 1000 employees 1000*4=4000mb = 4gb, will need to find effecient way to fetch this
*
 */

func (pc *payslipConfig) generatePaySlip(imgUrl string, quality int8) error {
	if quality < 0 || quality > 100 {
		return fmt.Errorf("Qualit should be in range 0-100 recived: %d", quality)
	}

	// load image
	res, err := http.Get(imgUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	bytesData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	pc.image = base64.StdEncoding.EncodeToString(bytesData)

	return nil
}
