package model

type TrainingMenu struct {
	ID          int64
	Name        string
	Description string
	Menu        []TrainingSet
}

type TrainingSet struct {
	ID          int64
	Name        string
	Description string
	ExerciseID  int64
	Reps        int64
	Weight      float64
	Interval    int64
}

type TrainingExercise struct {
	ID          int64
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
	Chest      = TargetMuscle("chest")
	Abdominals = TargetMuscle("abdominals")
	Quads      = TargetMuscle("quads")
	Biceps     = TargetMuscle("biceps")
	Shoulders  = TargetMuscle("shoulders")
	Lats       = TargetMuscle("lats")

	Stretches  = TrainingCategory("stretches")
	BodyWeight = TrainingCategory("body_weight")
	Barbell    = TrainingCategory("barbell")
	Dumbbells  = TrainingCategory("dumbbells")

	Beginner     = TrainingDifficulty("beginner")
	Intermediate = TrainingDifficulty("intermediate")
	Advanced     = TrainingDifficulty("advanced")
)
