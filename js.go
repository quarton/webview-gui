package gui

// webview-gui control methods avaible to Javascript

type GuiExport struct {
	gui *Gui
}

func (g *GuiExport) Terminate(args struct{}, reply *struct{}) error {
	g.gui.Terminate()
	return nil
}

func (g *GuiExport) SetTitle(title string, reply *struct{}) error {
	g.gui.SetTitle(title)
	return nil
}

func (g *GuiExport) SetFullScreen(fullScreen bool, reply *struct{}) error {
	g.gui.SetFullScreen(fullScreen)
	return nil
}

type ColorRGBArgs struct {
	R, G, B uint8
}

func (g *GuiExport) SetColorRGB(args ColorRGBArgs, reply *struct{}) error {
	g.gui.SetColorRGB(args.R, args.G, args.B)
	return nil
}

// File and directory dialogs

type PathReply struct {
	Path string
}

func (g *GuiExport) DialogOpenDirectory(title string, reply *PathReply) error {
	reply.Path = g.gui.DialogOpenDirectory(title)
	return nil
}

func (g *GuiExport) DialogOpenFile(title string, reply *PathReply) error {
	reply.Path = g.gui.DialogOpenFile(title)
	return nil
}

func (g *GuiExport) DialogSaveFile(title string, reply *PathReply) error {
	reply.Path = g.gui.DialogSaveFile(title)
	return nil
}

// Alerts

type AlertArgs struct {
	Title   string
	Message string
}

func (g *GuiExport) DialogAlertInfo(args *AlertArgs, reply *struct{}) error {
	g.gui.DialogAlertInfo(args.Title, args.Message)
	return nil
}

func (g *GuiExport) DialogAlertWarning(args *AlertArgs, reply *struct{}) error {
	g.gui.DialogAlertWarning(args.Title, args.Message)
	return nil
}

func (g *GuiExport) DialogAlertError(args *AlertArgs, reply *struct{}) error {
	g.gui.DialogAlertError(args.Title, args.Message)
	return nil
}
