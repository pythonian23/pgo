package flags

// Args is a type for holding variables from parsed FlagSets.
type Args struct {
	generalArgs
	identityArgs
	apiArgs
}

type generalArgs struct {
	Help *bool
}
type identityArgs struct {
	Alliance    *bool
	ID          *int
	Link        *string
	DiscordName *string
}
type apiArgs struct {
	APIKey *string
}
