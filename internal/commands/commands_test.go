package commands

import "testing"

func TestApp_Commands(t *testing.T) {
	t.Parallel()
	seen := map[string]bool{}
	for _, c := range (App{}).Commands() {
		if len(c.Names) == 0 {
			t.Errorf("comando sin nombres: %+v", c)
		}
		if c.Run == nil {
			t.Errorf("%v: Run es nil", c.Names)
		}
		if c.Description == "" {
			t.Errorf("%v: sin descripción", c.Names)
		}
		if c.MaxArgs != Unlimited && c.MaxArgs < c.MinArgs {
			t.Errorf("%v: MaxArgs %d < MinArgs %d", c.Names, c.MaxArgs, c.MinArgs)
		}
		if (c.MinArgs > 0 || c.MaxArgs != 0) && c.Arg == "" {
			t.Errorf("%v: acepta argumentos pero no los documenta en Arg", c.Names)
		}
		for _, n := range c.Names {
			if seen[n] {
				t.Errorf("nombre de comando duplicado: %s", n)
			}
			seen[n] = true
		}
	}
}
