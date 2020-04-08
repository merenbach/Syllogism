package symbol

// A Table of symbols.
type Table []*Symbol

// Dump values of variables in a symbol table.
// TODO: improve alignment?
// func (st Table) Dump() string {
// 	dump := new(bytes.Buffer)
// 	if len(st) > 0 {
// 		w := tabwriter.NewWriter(dump, 0, 0, 2, ' ', 0)
// 		fmt.Fprint(w, "Adr.\tart.\tterm\ttype\toccurs\tdist. count")
// 		for i, s := range st {
// 			fmt.Fprintf(w, "\n%d\t%s\t%d\t%d", i+1, s.Dump(), s.Occurrences, s.DistributionCount)
// 		}
// 		w.Flush()
// 	}
// 	return dump.String()
// }
