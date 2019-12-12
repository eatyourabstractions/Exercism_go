package house

var items = []struct {
	verb string
	clause string
}{
	{verb: "lay in ", clause: "the house that Jack built."},
	{verb: "ate " , clause: "the malt"},
	{verb: "killed ", clause: "the rat"},
	{verb: "worried ", clause: "the cat"},
	{verb: "tossed ", clause: "the dog"},
	{verb: "milked ", clause: "the cow with the crumpled horn"},
	{verb: "kissed ", clause: "the maiden all forlorn"},
	{verb: "married ", clause: "the man all tattered and torn"},
	{verb: "woke ", clause: "the priest all shaven and shorn"},
	{verb: "kept ", clause: "the rooster that crowed in the morn"},
	{verb: "belonged to ", clause: "the farmer sowing his corn"},
	{verb: "", clause: "the horse and the hound and the horn"},
}
func that(num int)(that_line string) {
	if num != -1 {
		that_line = "\nthat " + items[num].verb + items[num].clause +
				 that(num - 1)
	} else {
		that_line = ""
	}
	return
}
func this(num int)(this_line string) {
	this_line = "This is " + items[num].clause
	return
}

func Verse(n int)(v string) {
	n--
	if n < 1{
		v = this(n)
	} else{
		v = this(n) + that(n - 1)
	}
	return
}

func Song()(song string){
	for v := 1; v <= 11; v++{
		song += Verse(v) + "\n\n"
	}
	song += Verse(12)
	return
}