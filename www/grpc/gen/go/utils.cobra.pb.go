// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pactus

import (
	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	flag "github.com/NathanBaulch/protoc-gen-cobra/flag"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func UtilsClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Utils"),
		Short: "Utils service client",
		Long:  "Utils service defines RPC methods for utility functions such as message\n signing and verification.",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_UtilsSignMessageWithPrivateKeyCommand(cfg),
		_UtilsVerifyMessageCommand(cfg),
	)
	return cmd
}

func _UtilsSignMessageWithPrivateKeyCommand(cfg *client.Config) *cobra.Command {
	req := &SignMessageWithPrivateKeyRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("SignMessageWithPrivateKey"),
		Short: "SignMessageWithPrivateKey RPC client",
		Long:  "SignMessageWithPrivateKey sign message with provided private key.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Utils"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Utils", "SignMessageWithPrivateKey"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewUtilsClient(cc)
				v := &SignMessageWithPrivateKeyRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.SignMessageWithPrivateKey(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.PrivateKey, cfg.FlagNamer("PrivateKey"), "", "The private key to sign the message.")
	cmd.PersistentFlags().StringVar(&req.Message, cfg.FlagNamer("Message"), "", "The message to sign.")

	return cmd
}

func _UtilsVerifyMessageCommand(cfg *client.Config) *cobra.Command {
	req := &VerifyMessageRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("VerifyMessage"),
		Short: "VerifyMessage RPC client",
		Long:  "VerifyMessage verify signature with public key and message",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Utils"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Utils", "VerifyMessage"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewUtilsClient(cc)
				v := &VerifyMessageRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.VerifyMessage(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Message, cfg.FlagNamer("Message"), "", "The signed message.")
	cmd.PersistentFlags().StringVar(&req.Signature, cfg.FlagNamer("Signature"), "", "The signature of the message.")
	cmd.PersistentFlags().StringVar(&req.PublicKey, cfg.FlagNamer("PublicKey"), "", "The public key of the signer.")

	return cmd
}