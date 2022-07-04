package eapFact

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type FactEtab struct {
	Owner_civility string `db:"owner_civility"`
	Owner_name     string `db:"owner_name"`
	Owner_surname  string `db:"owner_surname"`
	Mail           string `db:"mail"`
	Phone          string `db:"phone"`
	Name           string `db:"name"`
	Fact_addr      string `db:"fact_addr"`
	Fact_cp        int    `db:"fact_cp"`
	Fact_city      string `db:"fact_city"`
	Fact_country   string `db:"fact_country"`
	Offer          int    `db:"offer"`
	Fact_infos     FactInfos
	Etab_offer     Offer
}

// Fact making structure
type Offer struct {
	Id       int     `db:"id"`
	Name     string  `db:"name"`
	PriceHT  float64 `db:"priceHT"`
	PriceTTC float64 `db:"priceTTC"`
}

type FactInfos struct {
	Id      int64
	Uuid    string `db:"uuid"`
	IsFirst bool
	Link    string `db:"link"`
	Date    string `db:"created"`
}

// Ticket making structures
type Order struct {
	Cli_uuid    string        `json:"cli_uuid"`
	Token       string        `json:"token"`
	TotalTTC    float64       `json:"totalTTC"`
	TotalHT     float64       `json:"totalHT"`
	Order_items []*OrderItems `json:"Order_items"`
}

type OrderItems struct {
	Item_id  int     `json:"item_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Infos struct {
	Name         string `db:"name"`
	Fact_addr    string `db:"addr"`
	Fact_cp      int    `db:"cp"`
	Fact_city    string `db:"city"`
	Fact_country string `db:"country"`
	Picture      string `db:"picture"`
}

func CreateFact(factInfos FactEtab) (err error) {

	m := pdf.NewMaroto(consts.Portrait, consts.Letter)

	m.Row(20, func() {
		m.Col(4, func() {
			m.Text(factInfos.Fact_infos.Date, props.Text{
				Top:         12,
				Size:        8,
				Extrapolate: true,
			})
		})
		m.ColSpace(6)
		m.Col(4, func() {
			m.Text("EasyAsPie - Facture "+strconv.FormatInt(factInfos.Fact_infos.Id, 10), props.Text{
				Top:         12,
				Size:        8,
				Extrapolate: true,
			})
		})
	})

	m.Row(50, func() {

		m.Col(4, func() {
			_ = m.FileImage("logo.png", props.Rect{
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
			m.Text(factInfos.Name, props.Text{
				Top:         10,
				Size:        11,
				Extrapolate: true,
			})
			m.Text(factInfos.Fact_addr, props.Text{
				Top:         15,
				Size:        8,
				Extrapolate: true,
			})
			m.Text(strconv.Itoa(factInfos.Fact_cp)+" "+factInfos.Fact_city, props.Text{
				Top:         18,
				Size:        8,
				Extrapolate: true,
			})
			m.Text(factInfos.Fact_country, props.Text{
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
			m.Text("Numéro de Facture: "+strconv.FormatInt(factInfos.Fact_infos.Id, 10), props.Text{
				Size: 10,
				Top:  2,
			})
		})
		m.ColSpace(2)
		m.Col(6, func() {
			m.Text("Date: "+factInfos.Fact_infos.Date, props.Text{
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
			m.Text("Montant TTC € ", props.Text{
				Size:  8,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	// Fact  Items
	m.Row(7, func() {
		m.Col(6, func() {
			m.Text(" Abonnement \""+factInfos.Etab_offer.Name+"\"", props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(fmt.Sprintf("%.2f", factInfos.Etab_offer.PriceHT), props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text(fmt.Sprintf("%.2f", factInfos.Etab_offer.PriceTTC), props.Text{
				Size:  8,
				Top:   1,
				Align: consts.Center,
			})
		})
	})

	totalHT := factInfos.Etab_offer.PriceHT
	totalTTC := factInfos.Etab_offer.PriceTTC

	if factInfos.Fact_infos.IsFirst {

		totalHT = factInfos.Etab_offer.PriceHT - factInfos.Etab_offer.PriceHT
		totalTTC = factInfos.Etab_offer.PriceTTC - factInfos.Etab_offer.PriceTTC

		m.Row(7, func() {
			m.Col(6, func() {
				m.Text(" Premier mois offert", props.Text{
					Size:  8,
					Top:   1,
					Align: consts.Left,
				})
			})
			m.Col(3, func() {
				m.Text("-"+fmt.Sprintf("%.2f", factInfos.Etab_offer.PriceHT), props.Text{
					Size:  8,
					Top:   1,
					Align: consts.Center})
			})
			m.Col(3, func() {
				m.Text("-"+fmt.Sprintf("%.2f", factInfos.Etab_offer.PriceTTC), props.Text{
					Size:  8,
					Top:   1,
					Align: consts.Center})
			})
		})
	}
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
			m.Text(fmt.Sprintf("%.2f", totalHT)+" €", props.Text{
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
			m.Text(fmt.Sprintf("%.2f", totalTTC)+" €", props.Text{
				Size:  15,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center})
		})
	})

	m.SetBorder(false)

	err = m.OutputFileAndClose(factInfos.Fact_infos.Link)
	if err != nil {
		fmt.Println("Could not save PDF:", err)
	}

	return err

}

func CreateTicket(id int64, dest string, PLOrder Order, etab Infos) (err error) {

	m := pdf.NewMaroto(consts.Portrait, consts.Letter)

	m.Row(20, func() {
		m.Col(4, func() {
			m.Text(time.Now().Format("2006-02-01 15:04:05"), props.Text{
				Top:         12,
				Size:        8,
				Extrapolate: true,
			})
		})
		m.ColSpace(6)
		m.Col(4, func() {
			m.Text("EasyAsPie - Commande n°"+strconv.FormatInt(id, 10), props.Text{
				Top:         12,
				Size:        8,
				Extrapolate: true,
			})
		})
	})

	m.Row(50, func() {

		m.Col(4, func() {
			_ = m.FileImage("/home/ec2-user/media/pictures/"+strings.SplitN(etab.Picture, "/", 5)[4], props.Rect{
				Left:    0,
				Top:     9,
				Percent: 50,
			})
			m.Text(etab.Name, props.Text{
				Top:         34,
				Size:        8,
				Extrapolate: true,
			})
			m.Text(etab.Fact_addr+", "+strconv.Itoa(etab.Fact_cp)+" "+etab.Fact_city, props.Text{
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
	})

	m.SetBorder(true)

	m.Row(7, func() {
		m.Col(4, func() {
			m.Text("Référence ", props.Text{
				Size:  8,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(2, func() {
			m.Text("Quantité ", props.Text{
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
			m.Text("Montant TTC € ", props.Text{
				Size:  8,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	for _, item := range PLOrder.Order_items {

		// Fact  Items
		m.Row(7, func() {
			m.Col(4, func() {
				m.Text("   "+item.Name, props.Text{
					Size:  8,
					Top:   1,
					Align: consts.Left,
				})
			})
			m.Col(2, func() {
				m.Text(strconv.Itoa(item.Quantity), props.Text{
					Size:  8,
					Top:   1,
					Align: consts.Center,
				})
			})
			m.Col(3, func() {
				m.Text(fmt.Sprintf("%.2f", item.Price/1.2), props.Text{
					Size:  8,
					Top:   1,
					Align: consts.Center,
				})
			})
			m.Col(3, func() {
				m.Text(fmt.Sprintf("%.2f", item.Price), props.Text{
					Size:  8,
					Top:   1,
					Align: consts.Center,
				})
			})
		})
	}

	totalHT := PLOrder.TotalHT
	totalTTC := PLOrder.TotalTTC

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
			m.Text(fmt.Sprintf("%.2f", totalHT)+" €", props.Text{
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
			m.Text("PAYÉ TTC :", props.Text{
				Size:  15,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center})
		})
		m.Col(3, func() {
			m.Text(fmt.Sprintf("%.2f", totalTTC)+" €", props.Text{
				Size:  15,
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center})
		})
	})

	m.SetBorder(false)

	err = m.OutputFileAndClose(dest)
	if err != nil {
		fmt.Println("Could not save PDF:", err)
	}

	return err

}
