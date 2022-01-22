package model

type TrainingExcercise struct {
	ID          string
	Name        string
	Description string
	Target      TargetMuscle
	Category    TrainingCategory
	Difficulty  TrainingDifficulty
}

type TargetMuscle string
type TrainingCategory string
type TrainingDifficulty string

const (
	Beginner     = TrainingDifficulty("beginner")
	Intermediate = TrainingDifficulty("intermediate")
	Advanced     = TrainingDifficulty("advanced")
)

const (
	Stretches  = TrainingCategory("stretches")
	BodyWeight = TrainingCategory("body_weight")
	Barbell    = TrainingCategory("barbell")
	Dumbbells  = TrainingCategory("dumbbells")
)

const (
	Chest      = TargetMuscle("chest")
	Abdominals = TargetMuscle("abdominals")
	Quads      = TargetMuscle("quads")
	Biceps     = TargetMuscle("biceps")
	Shoulders  = TargetMuscle("shoulders")
	Lats       = TargetMuscle("lats")
)
