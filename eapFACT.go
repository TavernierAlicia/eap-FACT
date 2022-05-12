package eapFact

import (
	"fmt"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func CreateFact() {
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)

	m.Row(20, func() {
		m.Col(4, func() {
			m.Text(time.Now().Format("2006-01-02 15:04:05"), props.Text{
				Top:         12,
				Size:        8,
				Extrapolate: true,
			})
		})
		m.ColSpace(6)
		m.Col(4, func() {
			m.Text("EasyAsPie - Facture 33", props.Text{
				Top:         12,
				Size:        8,
				Extrapolate: true,
			})
		})
	})

	m.Row(50, func() {

		m.Col(4, func() {
			_ = m.FileImage("./logo.png", props.Rect{
				Left:    0,
				Top:     9,
				Percent: 50,
			})
			m.Text("easy-as-pie.fr", props.Text{
				Top:         34,
				Size:        8,
				Extrapolate: true,
			})
			m.Text("20 rue de Flandres, 75019 Paris", props.Text{
				Top:         37,
				Size:        8,
				Extrapolate: true,
			})
			m.Text("", props.Text{
				Top:         47,
				Size:        8,
				Extrapolate: true,
			})

		})

		m.ColSpace(4)
		m.Col(4, func() {
			m.Text("Le Baraque o Bahamas SAS", props.Text{
				Top:         10,
				Size:        11,
				Extrapolate: true,
			})
			m.Text("10 rue des Lilas", props.Text{
				Top:         15,
				Size:        8,
				Extrapolate: true,
			})
			m.Text("75010 Paris", props.Text{
				Top:         18,
				Size:        8,
				Extrapolate: true,
			})
			m.Text("FRANCE", props.Text{
				Top:         21,
				Size:        8,
				Extrapolate: true,
			})
			m.Text("Paiement par prélèvements (RIB)", props.Text{
				Top:         30,
				Size:        6,
				Extrapolate: true,
			})
		})

	})

	m.Row(7, func() {
		m.Col(4, func() {
			m.Text("Numéro de Facture: 33", props.Text{
				Size: 10,
				Top:  2,
			})
		})
		m.ColSpace(2)
		m.Col(6, func() {
			m.Text("Date: "+time.Now().Format("2006-01-02 15:04:05"), props.Text{
				Size:  10,
				Top:   2,
				Align: consts.Right,
			})
		})
	})

	m.SetBorder(true)

	m.Row(7, func() {
		m.Col(6, func() {
			m.Text("Référence ", props.Text{
				Size:  8,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text("Prix HT € ", props.Text{
				Size:  8,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text("Montant € ", props.Text{
				Size:  8,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	// Fact Items
	m.Row(7, func() {
		m.Col(6, func() {
			m.Text(" Abonnement \"essential\"", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text("15,00", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text("15,00", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Center,
			})
		})
	})

	m.Row(7, func() {
		m.Col(6, func() {
			m.Text(" Premier mois offert", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text("-15,00", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Center})
		})
		m.Col(3, func() {
			m.Text("-15,00", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Center})
		})
	})

	m.Row(7, func() {})

	m.Row(7, func() {
		m.SetBorder(false)
		m.Col(6, func() {
			m.Text("", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Left})
		})
		m.SetBorder(true)
		m.Col(3, func() {
			m.Text("Total HT ", props.Text{
				Size:  8,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center})
		})
		m.Col(3, func() {
			m.Text("0,00€", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Center})
		})
	})

	m.Row(7, func() {
		m.SetBorder(false)
		m.Col(6, func() {
			m.Text("", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Left})
		})
		m.SetBorder(true)
		m.Col(3, func() {
			m.Text("TVA ", props.Text{
				Size:  8,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center})
		})
		m.Col(3, func() {
			m.Text("20%", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Center})
		})
	})

	m.Row(7, func() {})

	m.SetBorder(false)
	m.Row(7, func() {
		m.Col(6, func() {
			m.Text("", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Left})
		})
		m.Col(3, func() {
			m.Text("À PAYER TTC :", props.Text{
				Size:  15,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center})
		})
		m.Col(3, func() {
			m.Text("0,00 €", props.Text{
				Size:  15,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center})
		})
	})

	m.SetBorder(false)

	// generate name
	// name := time.Now().Format("2006-01-02") + "_"+".pdf"
	// fmt.Println(name)

	err := m.OutputFileAndClose("./facts/zpl.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

}
