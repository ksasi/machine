package commands

import (
	"errors"
	"fmt"

	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/log"
)

func cmdRm(c CommandLine, api libmachine.API) error {
	if len(c.Args()) == 0 {
		c.ShowHelp()
		return errors.New("You must specify a machine name")
	}

	force := c.Bool("force")

	for _, hostName := range c.Args() {
		h, err := api.Load(hostName)
		if err != nil {
			return fmt.Errorf("Error removing host %q: %s", hostName, err)
		}

		if err := h.Driver.Remove(); err != nil {
			if !force {
				log.Errorf("Provider error removing machine %q: %s", hostName, err)
				continue
			}
		}

		if err := api.Remove(hostName); err != nil {
			log.Errorf("Error removing machine %q from store: %s", hostName, err)
		} else {
			log.Infof("Successfully removed %s", hostName)
		}
	}

	return nil
}
