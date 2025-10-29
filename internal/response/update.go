package response

import (
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func (m ResponseModel) Init() tea.Cmd {
	return nil
}

func (m ResponseModel) Update(msg tea.Msg) (ResponseModel, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "down" {
			m.Viewport.GotoBottom()
		}
	case tea.WindowSizeMsg:
		// minus 1 for text header
		m.ResponseHeight = msg.Height*90/100 - utils.BoxStyle.GetVerticalBorderSize() - 1
		m.ResponseWidth = msg.Width*60/100 - utils.BoxStyle.GetHorizontalBorderSize()

		if !m.ViewportReady {

			m.Viewport = viewport.New(m.ResponseWidth, m.ResponseHeight)
			m.Viewport.YPosition = 0
			// m.Viewport.SetContent(string(m.Result.Data))
			m.Viewport.SetContent(Dummy)
			m.ViewportReady = true

		} else {
			m.Viewport.Width = m.ResponseWidth
			m.Viewport.Height = m.ResponseHeight
		}
	}
	return m, nil
}

var Dummy = `{
		// an ID field has been added here, however it's not required. You could use
		// any text field as long as it's unique for the zone.
		item{id: "item_1", title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{id: "item_2", title: "Nutella", desc: "It's good on toast"},
		item{id: "item_3", title: "Bitter melon", desc: "It cools you down"},
		item{id: "item_4", title: "Nice socks", desc: "And by that I mean socks without holes"},
		item{id: "item_5", title: "Eight hours of sleep", desc: "I had this once"},
		item{id: "item_6", title: "Cats", desc: "Usually"},
		item{id: "item_7", title: "Plantasia, the album", desc: "My plants love it too"},
		item{id: "item_8", title: "Pour over coffee", desc: "It takes forever to make though"},
		item{id: "item_9", title: "VR", desc: "Virtual reality...what is there to say?"},
		item{id: "item_10", title: "Noguchi Lamps", desc: "Such pleasing organic forms"},
		item{id: "item_11", title: "Linux", desc: "Pretty much the best OS"},
		item{id: "item_12", title: "Business school", desc: "Just kidding"},
		item{id: "item_13", title: "Pottery", desc: "Wet clay is a great feeling"},
		item{id: "item_14", title: "Shampoo", desc: "Nothing like clean hair"},
		item{id: "item_15", title: "Table tennis", desc: "It’s surprisingly exhausting"},
		item{id: "item_16", title: "Milk crates", desc: "Great for packing in your extra stuff"},
		item{id: "item_17", title: "Afternoon tea", desc: "Especially the tea sandwich part"},
		item{id: "item_18", title: "Stickers", desc: "The thicker the vinyl the better"},
		item{id: "item_19", title: "20° Weather", desc: "Celsius, not Fahrenheit"},
		item{id: "item_20", title: "Warm light", desc: "Like around 2700 Kelvin"},
		item{id: "item_21", title: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
		item{id: "item_22", title: "Gaffer’s tape", desc: "Basically sticky fabric"},
		item{id: "item_23", title: "Terrycloth", desc: "In other words, towel fabric"},
	}`
