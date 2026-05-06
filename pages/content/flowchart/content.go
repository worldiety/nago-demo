package flowchart

import (
	_ "embed"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/flowchart"
)

//go:embed joe_schmoe_green.svg
var JoeSchmoeGreen core.SVG

//go:embed joe_schmoe_red.svg
var JoeSchmoeRed core.SVG

func Content(wnd core.Window) core.View {
	personNodes := make([]flowchart.Node, 0)
	personContents := make([]flowchart.CustomContent, 0)

	jbusseNode, jbusseContent := personNode("jochen-busse", "Jochen Busse", "Geschäftsführer", JoeSchmoeGreen, flowchart.Point{
		X: 100,
		Y: -50,
	}, flowchart.NodeTypeStart)
	personNodes = append(personNodes, jbusseNode)
	personContents = append(personContents, jbusseContent)

	gkoesterNode, gkoesterContent := personNode("gabi-koester", "Gabi Köster", "Projektmanager", JoeSchmoeGreen, flowchart.Point{
		X: -100,
		Y: 100,
	}, flowchart.NodeTypeDefault)
	personNodes = append(personNodes, gkoesterNode)
	personContents = append(personContents, gkoesterContent)

	kpohlNode, kpohlContent := personNode("kalle-pohl", "Kalle Pohl", "Projektmanager", JoeSchmoeGreen, flowchart.Point{
		X: 300,
		Y: 100,
	}, flowchart.NodeTypeDefault)
	personNodes = append(personNodes, kpohlNode)
	personContents = append(personContents, kpohlContent)

	gcantzNode, gcantzContent := personNode("guido-cantz", "Guido Cantz", "Projektmanager", JoeSchmoeGreen, flowchart.Point{
		X: 100,
		Y: 100,
	}, flowchart.NodeTypeDefault)
	personNodes = append(personNodes, gcantzNode)
	personContents = append(personContents, gcantzContent)

	bstelterNode, bstelterContent := personNode("bernd-stelter", "Bernd Stelter", "Softwareentwickler", JoeSchmoeRed, flowchart.Point{
		X: -100,
		Y: 250,
	}, flowchart.NodeTypeEnd)
	personNodes = append(personNodes, bstelterNode)
	personContents = append(personContents, bstelterContent)

	gcantznichtNode, gcantznichtContent := personNode("guido-cantz-nicht", "Guido Cantz Nicht", "Praktikant", JoeSchmoeGreen, flowchart.Point{
		X: 100,
		Y: 250,
	}, flowchart.NodeTypeEnd)
	personNodes = append(personNodes, gcantznichtNode)
	personContents = append(personContents, gcantznichtContent)

	edges := []flowchart.Edge{
		{
			ID:           "jbusse-gkoester",
			SourceNodeID: "jochen-busse",
			TargetNodeID: "gabi-koester",
			Animated:     true,
		},
		{
			ID:           "jbusse-kpohl",
			SourceNodeID: "jochen-busse",
			TargetNodeID: "kalle-pohl",
			Animated:     true,
		},
		{
			ID:           "jbusse-gcantz",
			SourceNodeID: "jochen-busse",
			TargetNodeID: "guido-cantz",
			Animated:     true,
		},
		{
			ID:           "gkoester-bstelter",
			SourceNodeID: "gabi-koester",
			TargetNodeID: "bernd-stelter",
			Label:        "hat entlassen",
			Width:        2,
			Style:        flowchart.EdgeStyleDashed,
			Color:        ui.ColorError,
			MarkerEnd:    flowchart.EdgeMarkerArrow,
		},
		{
			ID:           "gcantz-gcantznicht",
			SourceNodeID: "guido-cantz",
			TargetNodeID: "guido-cantz-nicht",
		},
	}

	state := core.AutoState[flowchart.Model](wnd).Init(func() flowchart.Model {
		return flowchart.Model{
			Nodes: personNodes,
			Edges: edges,
		}
	})

	return ui.Stack(
		flowchart.FlowChart(state.Get()).
			InputValue(state).
			NodesDraggable(true).
			NodesConnectable(true).
			ElementsSelectable(true).
			BackgroundColor(ui.M1).
			Frame(ui.Frame{}.FullWidth().FullHeight()).
			Layout(flowchart.FlowChartLayoutVertical).
			CustomContents(personContents).
			MaxZoom(1.5).
			BackgroundColor(ui.M2.WithTransparency(65)),
	).
		Frame(ui.Frame{Height: ui.L560}.FullWidth()).
		Border(ui.Border{}.Radius(ui.L14))
}

func personNode(id, name, title string, icon core.SVG, position flowchart.Point, nodeType flowchart.NodeType) (flowchart.Node, flowchart.CustomContent) {
	return flowchart.Node{
			ID:       id,
			Position: position,
			Label:    name,
			Type:     nodeType,
		},
		flowchart.CustomContent{
			NodeID: id,
			Content: ui.VStack(
				ui.ImageIcon(icon),
				ui.VStack(
					ui.Text(name).Font(ui.TitleLarge),
					ui.Text(title).Font(ui.TitleSmall),
				),
			).Gap(ui.L4).Padding(ui.Padding{}.All(ui.L8)),
		}
}
