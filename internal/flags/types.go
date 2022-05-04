package flags

type Args struct {
	*generalArgs
	*identityArgs
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
