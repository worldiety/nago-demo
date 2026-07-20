package dataview

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/worldiety/option"
	"go.wdy.de/nago/pkg/xtime"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
	"go.wdy.de/nago/presentation/ui/dataview"
	"go.wdy.de/nago/presentation/ui/pager"
)

type (
	VID          string
	VBrand       string
	VModel       string
	VHorsePower  int64
	VReleaseDate xtime.Date
)

type Vehicle struct {
	ID          VID
	Brand       VBrand
	Model       VModel
	HorsePower  VHorsePower
	ReleaseDate VReleaseDate
}

func (v Vehicle) Identity() VID {
	return v.ID
}

func (v Vehicle) String() string {
	return fmt.Sprintf("%s %s", v.Brand, v.Model)
}

var vehicles = []Vehicle{
	{ID: "bmw-m3", Brand: "BMW", Model: "M3", HorsePower: 400, ReleaseDate: VReleaseDate(xtime.DateNow())},
	{ID: "audi-rs7", Brand: "Audi", Model: "RS7", HorsePower: 500, ReleaseDate: VReleaseDate(xtime.Date{Year: 2021, Month: 5, Day: 15})},
	{ID: "mercedes-amg-gt", Brand: "Mercedes", Model: "AMG GT", HorsePower: 600, ReleaseDate: VReleaseDate(xtime.Date{Year: 2021, Month: 5, Day: 15})},
	{ID: "porsche-911", Brand: "Porsche", Model: "911", HorsePower: 700, ReleaseDate: VReleaseDate(xtime.Date{Year: 2023, Month: 2, Day: 1})},
	{ID: "ferrari-488", Brand: "Ferrari", Model: "488", HorsePower: 800, ReleaseDate: VReleaseDate(xtime.Date{Year: 2024, Month: 2, Day: 2})},
	{ID: "fiat-punto", Brand: "Fiat", Model: "Punto", HorsePower: 69, ReleaseDate: VReleaseDate(xtime.Date{Year: 2005, Month: 11, Day: 27})},
	{ID: "ford-mustang", Brand: "Ford", Model: "Mustang", HorsePower: 450, ReleaseDate: VReleaseDate(xtime.Date{Year: 2019, Month: 5, Day: 15})},
	{ID: "chevrolet-camaro", Brand: "Chevrolet", Model: "Camaro", HorsePower: 455, ReleaseDate: VReleaseDate(xtime.Date{Year: 2018, Month: 3, Day: 10})},
	{ID: "nissan-gt-r", Brand: "Nissan", Model: "GT-R", HorsePower: 565, ReleaseDate: VReleaseDate(xtime.Date{Year: 2021, Month: 6, Day: 5})},
	{ID: "lamborghini-huracan", Brand: "Lamborghini", Model: "Huracán", HorsePower: 630, ReleaseDate: VReleaseDate(xtime.Date{Year: 2022, Month: 9, Day: 12})},
	{ID: "tesla-model-s", Brand: "Tesla", Model: "Model S", HorsePower: 1020, ReleaseDate: VReleaseDate(xtime.Date{Year: 2021, Month: 1, Day: 15})},
	{ID: "honda-civic", Brand: "Honda", Model: "Civic", HorsePower: 158, ReleaseDate: VReleaseDate(xtime.Date{Year: 2020, Month: 7, Day: 20})},
	{ID: "toyota-corolla", Brand: "Toyota", Model: "Corolla", HorsePower: 139, ReleaseDate: VReleaseDate(xtime.Date{Year: 2019, Month: 4, Day: 10})},
	{ID: "mazda-mx-5", Brand: "Mazda", Model: "MX-5", HorsePower: 181, ReleaseDate: VReleaseDate(xtime.Date{Year: 2021, Month: 3, Day: 5})},
}

func Content(wnd core.Window) core.View {
	createDialogPresented := core.AutoState[bool](wnd)

	return ui.Stack(
		dataview.FromData(wnd, dataview.Data[Vehicle, VID]{
			FindAll: func(yield func(VID, error) bool) {
				for _, v := range vehicles {
					if !yield(v.ID, nil) {
						return
					}
				}
			},
			FindByID: func(id VID) (option.Opt[Vehicle], error) {
				for _, v := range vehicles {
					if v.ID == id {
						return option.Some(Vehicle{ID: v.ID, Brand: v.Brand, Model: v.Model, HorsePower: v.HorsePower, ReleaseDate: v.ReleaseDate}), nil
					}
				}
				return option.None[Vehicle](), nil
			},
			Fields: []dataview.Field[Vehicle]{
				{
					ID:   "brand",
					Name: "Marke",
					Map: func(v Vehicle) core.View {
						return ui.Text(string(v.Brand))
					},
					Comparator: func(a, b Vehicle) int {
						return strings.Compare(strings.ToLower(string(a.Brand)), strings.ToLower(string(b.Brand)))
					},
				},
				{
					ID:   "model",
					Name: "Modell",
					Map: func(v Vehicle) core.View {
						return ui.Text(string(v.Model))
					},
					Comparator: func(a, b Vehicle) int {
						return strings.Compare(strings.ToLower(string(a.Model)), strings.ToLower(string(b.Model)))
					},
				},
				{
					ID:   "horsepower",
					Name: "Leistung (PS)",
					Map: func(v Vehicle) core.View {
						return ui.Text(fmt.Sprintf("%d", v.HorsePower))
					},
					Comparator: func(a, b Vehicle) int {
						return int(a.HorsePower - b.HorsePower)
					},
				},
				{
					ID:   "releaseDate",
					Name: "Erscheinungsdatum",
					Map: func(v Vehicle) core.View {
						return ui.Text(xtime.Date(v.ReleaseDate).Format("02.01.2006"))
					},
					Comparator: func(a, b Vehicle) int {
						if xtime.Date(a.ReleaseDate).After(xtime.Date(b.ReleaseDate)) {
							return -1
						}
						return 1
					},
				},
			},
		}).
			NextActionIndicator(true).
			SelectOptions(dataview.NewSelectOptionDelete(wnd, func(selected []VID) error {
				for _, id := range selected {
					vehicles = slices.DeleteFunc(vehicles, func(v Vehicle) bool {
						return v.ID == id
					})
				}
				return nil
			})).
			CreateAction(func() { createDialogPresented.Set(true) }).
			ModelOptions(pager.ModelOptions{PageSize: 10}).
			Search(true),
		createDialog(wnd, createDialogPresented),
	).Gap(ui.L32).FullWidth()
}

func createDialog(wnd core.Window, presented *core.State[bool]) core.View {
	if !presented.Get() {
		return nil
	}

	stateBrand := core.AutoState[string](wnd)
	stateModel := core.AutoState[string](wnd)
	stateHorsePower := core.AutoState[int64](wnd)
	stateReleaseDate := core.AutoState[xtime.Date](wnd)

	stateErrorBrand := core.AutoState[string](wnd)
	stateErrorModel := core.AutoState[string](wnd)
	stateErrorHorsePower := core.AutoState[string](wnd)
	stateErrorReleaseDate := core.AutoState[string](wnd)

	return alert.Dialog(
		"Neues Fahrzeug",
		ui.VStack(
			ui.TextField("Marke*", stateBrand.Get()).InputValue(stateBrand).FullWidth().ErrorText(stateErrorBrand.Get()),
			ui.TextField("Modell*", stateModel.Get()).InputValue(stateModel).FullWidth().ErrorText(stateErrorModel.Get()),
			ui.IntField("Leistung (PS)*", stateHorsePower.Get(), stateHorsePower).FullWidth().ErrorText(stateErrorHorsePower.Get()),
			ui.SingleDatePicker("Erscheinungsdatum*", stateReleaseDate.Get(), stateReleaseDate).ErrorText(stateErrorReleaseDate.Get()).Frame(ui.Frame{}.FullWidth()),
		).Gap(ui.L16).FullWidth(),
		presented,
		alert.Closeable(),
		alert.Create(func() (close bool) {
			brand := stateBrand.Get()
			if brand == "" {
				stateErrorBrand.Set("Marke darf nicht leer sein")
			} else {
				stateErrorBrand.Set("")
			}

			model := stateModel.Get()
			if model == "" {
				stateErrorModel.Set("Modell darf nicht leer sein")
			} else {
				stateErrorModel.Set("")
			}

			hp := stateHorsePower.Get()
			if hp <= 0 {
				stateErrorHorsePower.Set("Leistung muss größer als 0 sein")
			} else {
				stateErrorHorsePower.Set("")
			}

			releaseDate := stateReleaseDate.Get()
			if releaseDate.IsZero() {
				stateErrorReleaseDate.Set("Erscheinungsdatum darf nicht leer sein")
			} else {
				stateErrorReleaseDate.Set("")
			}

			if stateErrorBrand.Get() != "" || stateErrorModel.Get() != "" || stateErrorHorsePower.Get() != "" || stateErrorReleaseDate.Get() != "" {
				return false
			}

			vehicles = append(vehicles, Vehicle{
				ID:          VID(fmt.Sprintf("%s-%s", brand, model)),
				Brand:       VBrand(brand),
				Model:       VModel(model),
				HorsePower:  VHorsePower(hp),
				ReleaseDate: VReleaseDate(releaseDate),
			})

			stateBrand.Set("")
			stateModel.Set("")
			stateHorsePower.Set(0)
			stateReleaseDate.Set(xtime.Date{})

			stateErrorBrand.Set("")
			stateErrorModel.Set("")
			stateErrorHorsePower.Set("")
			stateErrorReleaseDate.Set("")

			return true
		}),
	)
}
