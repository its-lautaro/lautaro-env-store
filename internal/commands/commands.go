package commands

func (a App) Commands() []Command {
	return []Command{
		{Names: []string{"--list", "ls"}, Description: "Lista los nombres guardados y su fecha (nunca los valores) y los copia al clipboard", Run: a.List},
		{Names: []string{"--remove", "rm"}, Arg: "NOMBRE...", MinArgs: 1, MaxArgs: Unlimited, Description: "Borra esas variables", Run: a.Remove},
		{Names: []string{"--clear"}, Arg: "[--yes]", MaxArgs: 1, Description: "Borra todas; pide confirmación salvo --yes o -y", Run: a.Clear},
		{Names: []string{"--init"}, Description: "Crea el store y agrega el source al rc del shell", Run: a.Init},
		{Names: []string{"--help", "-h"}, Description: "Muestra esta ayuda", Run: a.Help},
	}
}
