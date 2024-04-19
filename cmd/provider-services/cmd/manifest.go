package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/akash-network/node/sdl"

	aclient "github.com/akash-network/provider/client"
	gwrest "github.com/akash-network/provider/gateway/rest"
)

var (
	errSubmitManifestFailed = errors.New("submit manifest to some providers has been failed")
)

// SendManifestCmd looks up the Providers blockchain information,
// and POSTs the SDL file to the Gateway address.
func SendManifestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "send-manifest <sdl-path>",
		Args:         cobra.ExactArgs(1),
		Short:        "Submit manifest to provider(s)",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doSendManifest(cmd, args[0])
		},
	}

	addManifestFlags(cmd)

	cmd.Flags().StringP(flagOutput, "o", outputText, "output format text|json|yaml. default text")

	return cmd
}

func doSendManifest(cmd *cobra.Command, sdlpath string) error {
	cctx, err := sdkclient.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	ctx := cmd.Context()

	cl, err := aclient.DiscoverQueryClient(ctx, cctx)
	if err != nil {
		return err
	}

	sdl, err := sdl.ReadFile(sdlpath)
	if err != nil {
		return err
	}

	mani, err := sdl.Manifest()
	if err != nil {
		return err
	}

	// cert, err := cutils.LoadAndQueryCertificateForAccount(cmd.Context(), cctx, nil)
	// if err != nil {
	// 	return markRPCServerError(err)
	// }

	dseq, err := dseqFromFlags(cmd.Flags())
	if err != nil {
		return err
	}

	// owner address in FlagFrom has already been validated thus save to just pull its value as string
	// leases, err := leasesForDeployment(cmd.Context(), cl, cmd.Flags(), dtypes.DeploymentID{
	// 	Owner: cctx.GetFromAddress().String(),
	// 	DSeq:  dseq,
	// })
	// if err != nil {
	// 	return markRPCServerError(err)
	// }

	// type result struct {
	// 	Provider     sdk.Address `json:"provider" yaml:"provider"`
	// 	Status       string      `json:"status" yaml:"status"`
	// 	Error        string      `json:"error,omitempty" yaml:"error,omitempty"`
	// 	ErrorMessage string      `json:"errorMessage,omitempty" yaml:"errorMessage,omitempty"`
	// }

	// results := make([]result, len(leases))
	// results := make([]result, 1)

	submitFailed := false

	prov, _ := sdk.AccAddressFromBech32("akash1vcgdh56ujtym8umkw3hj028aqu892qydsralwp")
	gclient, err := gwrest.NewClientILIJA(cl, prov)
	if err != nil {
		return err
	}
	err = gclient.SubmitManifest(cmd.Context(), dseq, mani)

	// ILIJA FIX 1
	// for i, lid := range leases {
	// 	prov, _ := sdk.AccAddressFromBech32(lid.Provider)
	// 	gclient, err := gwrest.NewClient(cl, prov, []tls.Certificate{cert})
	// 	if err != nil {
	// 		return err
	// 	}

	// 	err = gclient.SubmitManifest(cmd.Context(), dseq, mani)
	// 	res := result{
	// 		Provider: prov,
	// 		Status:   "PASS",
	// 	}
	// 	if err != nil {
	// 		res.Error = err.Error()
	// 		if e, valid := err.(gwrest.ClientResponseError); valid {
	// 			res.ErrorMessage = e.Message
	// 		}
	// 		res.Status = "FAIL"
	// 		submitFailed = true
	// 	}

	// 	results[i] = res
	// }

	// buf := &bytes.Buffer{}

	// switch cmd.Flag(flagOutput).Value.String() {
	// case outputText:
	// 	for _, res := range results {
	// 		_, _ = fmt.Fprintf(buf, "provider: %s\n\tstatus:       %s\n", res.Provider, res.Status)
	// 		if res.Error != "" {
	// 			_, _ = fmt.Fprintf(buf, "\terror:        %v\n", res.Error)
	// 		}
	// 		if res.ErrorMessage != "" {
	// 			_, _ = fmt.Fprintf(buf, "\terrorMessage: %v\n", res.ErrorMessage)
	// 		}
	// 	}
	// case outputJSON:
	// 	err = json.NewEncoder(buf).Encode(results)
	// case outputYAML:
	// 	err = yaml.NewEncoder(buf).Encode(results)
	// }

	if err != nil {
		return err
	}

	_, err = fmt.Println("Done WITH MANIFEST")

	if err != nil {
		return err
	}

	if submitFailed {
		return errSubmitManifestFailed
	}

	return nil
}
