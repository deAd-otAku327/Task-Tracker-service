package enum

type TaskRelation int

const (
	CreatedByMe TaskRelation = iota
	AssignedToMe
)

var relationToName = map[TaskRelation]string{
	CreatedByMe:  "created_by_me",
	AssignedToMe: "assigned_to_me",
}

var nameToRelation = map[string]TaskRelation{
	"created_by_me":  CreatedByMe,
	"assigned_to_me": AssignedToMe,
}

func CheckTaskRelation(name string) bool {
	_, ok := nameToRelation[name]
	return ok
}

func (s TaskRelation) String() string {
	return relationToName[s]
}
