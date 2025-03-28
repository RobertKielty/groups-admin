package cmd

// default command is ls - generate a report of lists and the role where profile is subscribed
// Flags
// 		--profile optional, defaults to configured user
//		--filter optional, default to, "" which will use all of configured user's lists
//      --output optional, default to one subscription per line, can select csv, json, yaml??

// sub-command add - add a new member to your subscriptions
// Flags
//		--profile, mandatory value [email addr] [no default], the team member that will be subscribed to your lists
//		--role, optional, value [owner|mod|member], their role when subscribed, defaults to Owner
//		--filter, optional, value [regular expression], used to restrict the lists to which they will be subscribed

// sub-command rm - remove an existing member from your lists
// Flags
//		--profile, mandatory value [email addr] [no default], the team member that will be removed from your lists
//		--filter, optional, value [regular expression], used to restrict the lists they will be removed from
