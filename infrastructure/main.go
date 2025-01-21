package main

import (
	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		repository, err := github.NewRepository(ctx, "tings", &github.RepositoryArgs{
			AllowMergeCommit:    pulumi.Bool(false),
			AllowRebaseMerge:    pulumi.Bool(true),
			AllowSquashMerge:    pulumi.Bool(false),
			AllowUpdateBranch:   pulumi.Bool(true),
			ArchiveOnDestroy:    pulumi.Bool(true),
			DeleteBranchOnMerge: pulumi.Bool(true),
			Description:         pulumi.String("A webapp for tracking personal data"),
			HasDiscussions:      pulumi.Bool(false),
			HasDownloads:        pulumi.Bool(false),
			HasIssues:           pulumi.Bool(false),
			HasProjects:         pulumi.Bool(false),
			HasWiki:             pulumi.Bool(false),
			LicenseTemplate:     pulumi.String("gpl-3.0"),
			Name:                pulumi.String("tings"),
			Topics: pulumi.StringArray{
				pulumi.String("go"),
			},
			Visibility:          pulumi.String("public"),
			VulnerabilityAlerts: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		ctx.Export("REPOSITORY_URL", repository.HttpCloneUrl)

		return nil
	})
}
