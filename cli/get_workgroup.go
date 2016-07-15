package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getWorkgroupHelp = `
workgroup [workgroupName]
View detailed workgroup information.
Examples:

	$ steam get workgroup production
`

func getWorkgroup(c *context) *cobra.Command {
	cmd := newCmd(c, getWorkgroupHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help get workgroup'.")
		}

		// -- Args --

		workgroupName := args[0]

		// -- Execution --

		workgroup, err := c.remote.GetWorkgroupByName(workgroupName)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		idents, err := c.remote.GetIdentitiesForWorkgroup(workgroup.Id)
		if err != nil {
			log.Fatalln(err)
		}

		// -- Formatting --

		base := []string{
			fmt.Sprintf("DESCRIPTION:\t%s", workgroup.Description),
			fmt.Sprintf("ID:\t%d", workgroup.Id),
			fmt.Sprintf("AGE:\t%s", fmtAgo(workgroup.Created)),
		}
		c.printt("\t"+workgroup.Name, base)

		fmt.Println("IDENTITIES:", len(idents))
		if len(idents) > 0 {
			ids := make([]string, len(idents))
			for i, id := range idents {
				ids[i] = fmt.Sprintf("%s\t%s\t%s",
					id.Name,
					identityStatus(id.IsActive),
					fmtAgo(id.LastLogin),
				)
			}
			c.printt("NAME\tSTATUS\tLAST LOGIN", ids)
		}

	})

	return cmd
}
