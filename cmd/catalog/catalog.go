package main

import (
	"fmt"
	"log"

	"github.com/caseymrm/menuet"
)

var alertsCatalog = []menuet.Alert{
	menuet.Alert{
		MessageText: "Just MessageText",
	},
	menuet.Alert{
		InformativeText: "Just InformativeText",
	},
	menuet.Alert{
		MessageText:     "MessageText and InformativeText",
		InformativeText: "This is the InformativeText",
	},
	menuet.Alert{
		MessageText: "Message and two buttons",
		Buttons:     []string{"One", "Two"},
	},
	menuet.Alert{
		MessageText: "Message and input",
		Inputs:      []string{"Example input"},
	},
	menuet.Alert{
		MessageText:     "Message, InformativeText, Button, and Input",
		InformativeText: "Example InformativeText",
		Buttons:         []string{"Example button"},
		Inputs:          []string{"Example Input"},
	},
	menuet.Alert{
		MessageText: "Message and two inputs",
		Inputs:      []string{"Input one", "Input two"},
	},
}

var notificationsCatalog = []menuet.Notification{
	menuet.Notification{
		Title: "Just a title",
	},
	menuet.Notification{
		Subtitle: "Just a subtitle",
	},
	menuet.Notification{
		Message: "Just a message",
	},
	menuet.Notification{
		Title:    "Title and subtitle",
		Subtitle: "This is the subtitle",
	},
	menuet.Notification{
		Title:   "Title and message",
		Message: "This is the message",
	},
	menuet.Notification{
		Subtitle: "Subtitle and message (this is the subtitle)",
		Message:  "This is the message",
	},
	menuet.Notification{
		Title:    "Title, subtitle, and message",
		Subtitle: "This is the subtitle",
		Message:  "This is the message",
	},
	menuet.Notification{
		Title:        "Action button",
		Subtitle:     "This is a subtitle",
		ActionButton: "Do an action",
	},
	menuet.Notification{
		Title:       "Close button",
		Subtitle:    "This is a subtitle",
		CloseButton: "Custom close button",
	},
	menuet.Notification{
		Title:               "ResponsePlaceholder ",
		Subtitle:            "This is a subtitle",
		ResponsePlaceholder: "Custom responsePlaceholder",
	},
	menuet.Notification{
		Title:      "Identifier set",
		Identifier: "identified",
	},
	menuet.Notification{
		Title: "Remove from notification center",
		RemoveFromNotificationCenter: true,
	},
}

func menuItems() []menuet.MenuItem {
	alerts := make([]menuet.MenuItem, 0, len(alertsCatalog))
	for ind, alert := range alertsCatalog {
		text := alert.MessageText
		if text == "" {
			text = alert.InformativeText
		}
		alerts = append(alerts, menuet.MenuItem{
			Text:     text,
			Callback: fmt.Sprintf("alert %d", ind),
		})
	}

	notifs := make([]menuet.MenuItem, 0, len(notificationsCatalog))
	for ind, notif := range notificationsCatalog {
		text := notif.Title
		if text == "" {
			text = notif.Subtitle
		}
		if text == "" {
			text = notif.Message
		}
		notifs = append(notifs, menuet.MenuItem{
			Text:     text,
			Callback: fmt.Sprintf("notif %d", ind),
		})
	}

	return []menuet.MenuItem{
		menuet.MenuItem{
			Text:     "Show Alert",
			Children: alerts,
		},
		menuet.MenuItem{
			Text:     "Send Notification",
			Children: notifs,
		},
		menuet.MenuItem{
			Text: "Menu Items",
			Children: []menuet.MenuItem{
				{
					Text: "Text no callback",
				},
				{
					Text:     "Text with Callback",
					Callback: "Text with Callback",
				},
				{
					Text: "FontSizes",
					Children: []menuet.MenuItem{
						{
							Text:     "FontSize 2",
							Callback: "FontSize 2",
							FontSize: 2,
						},
						{
							Text:     "FontSize 4",
							Callback: "FontSize 4",
							FontSize: 4,
						},
						{
							Text:     "FontSize 6",
							Callback: "FontSize 6",
							FontSize: 6,
						},
						{
							Text:     "FontSize 8",
							Callback: "FontSize 8",
							FontSize: 8,
						},
						{
							Text:     "FontSize 10",
							Callback: "FontSize 10",
							FontSize: 10,
						},
						{
							Text:     "FontSize 12",
							Callback: "FontSize 12",
							FontSize: 12,
						},
						{
							Text:     "FontSize 14",
							Callback: "FontSize 14",
							FontSize: 14,
						},
						{
							Text:     "FontSize 16",
							Callback: "FontSize 16",
							FontSize: 16,
						},
						{
							Text:     "FontSize 18",
							Callback: "FontSize 18",
							FontSize: 18,
						},
						{
							Text:     "FontSize 20",
							Callback: "FontSize 20",
							FontSize: 20,
						},
						{
							Text:     "FontSize 22",
							Callback: "FontSize 22",
							FontSize: 22,
						},
						{
							Text:     "FontSize 24",
							Callback: "FontSize 24",
							FontSize: 24,
						},
						{
							Text:     "FontSize 26",
							Callback: "FontSize 26",
							FontSize: 26,
						},
					},
				},
				{
					Text: "FontWeights",
					Children: []menuet.MenuItem{
						{
							Text:       "WeightUltraLight",
							FontWeight: menuet.WeightUltraLight,
							Callback:   "WeightUltraLight",
						},
						{
							Text:       "WeightThin",
							FontWeight: menuet.WeightThin,
							Callback:   "WeightThin",
						},
						{
							Text:       "WeightLight",
							FontWeight: menuet.WeightLight,
							Callback:   "WeightLight",
						},
						{
							Text:       "WeightRegular",
							FontWeight: menuet.WeightRegular,
							Callback:   "WeightRegular",
						},
						{
							Text:       "WeightMedium",
							FontWeight: menuet.WeightMedium,
							Callback:   "WeightMedium",
						},
						{
							Text:       "WeightSemibold",
							FontWeight: menuet.WeightSemibold,
							Callback:   "WeightSemibold",
						},
						{
							Text:       "WeightBold",
							FontWeight: menuet.WeightBold,
							Callback:   "WeightBold",
						},
						{
							Text:       "WeightHeavy",
							FontWeight: menuet.WeightHeavy,
							Callback:   "WeightHeavy",
						},
						{
							Text:       "WeightBlack",
							FontWeight: menuet.WeightBlack,
							Callback:   "WeightBlack",
						},
					},
				},
				{
					Text:     "State = true",
					Callback: "State = true",
					State:    true,
				},
			},
		},
	}
}

func handleClick(click string) {
	var index int
	var kind string
	n, err := fmt.Sscan(click, &kind, &index)
	if err != nil {
		log.Printf("Sscanf error: %v", err)
		return
	}
	if n != 2 {
		return
	}
	switch kind {
	case "notif":
		menuet.App().Notification(notificationsCatalog[index])
	case "alert":
		menuet.App().Alert(alertsCatalog[index])
	}
}

func main() {
	menuet.App().SetMenuState(&menuet.MenuState{
		Title: "Catalog",
		Items: menuItems(),
	})
	menuet.App().Label = "com.github.caseymrm.menuet.catalog"
	menuet.App().Clicked = handleClick
	menuet.App().RunApplication()
}