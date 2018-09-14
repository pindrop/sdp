package animation

// Code generated by cdproto-gen. DO NOT EDIT.

// EventAnimationCanceled event for when an animation has been cancelled.
type EventAnimationCanceled struct {
	ID string `json:"id"` // Id of the animation that was cancelled.
}

// EventAnimationCreated event for each animation that has been created.
type EventAnimationCreated struct {
	ID string `json:"id"` // Id of the animation that was created.
}

// EventAnimationStarted event for animation that has been started.
type EventAnimationStarted struct {
	Animation *Animation `json:"animation"` // Animation that was started.
}
