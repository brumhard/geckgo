package service

import (
	geckgov1 "github.com/brumhard/geckgo/pkg/pb/geckgo/v1"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MarshalListSettings(t *ListSettings) *geckgov1.ListSettings {
	if t == nil {
		return nil
	}

	if t.DailyTime == nil {
		return nil
	}

	return &geckgov1.ListSettings{DailyTime: durationpb.New(t.DailyTime.Duration)}
}

func UnmarshalListSettings(t *geckgov1.ListSettings) *ListSettings {
	if t == nil {
		return nil
	}

	if t.DailyTime == nil {
		return nil
	}

	return &ListSettings{DailyTime: &Duration{t.DailyTime.AsDuration()}}
}

func MarshalList(t *List) *geckgov1.List {
	if t == nil {
		return nil
	}

	return &geckgov1.List{
		Id:       int32(t.ID),
		Name:     t.Name,
		Settings: MarshalListSettings(t.Settings),
	}
}

func UnmarshalList(t *geckgov1.List) *List {
	if t == nil {
		return nil
	}

	return &List{
		ID:       int(t.Id),
		Name:     t.Name,
		Settings: UnmarshalListSettings(t.Settings),
	}
}

func MarshalMoment(t *Moment) *geckgov1.Moment {
	if t == nil {
		return nil
	}

	return &geckgov1.Moment{
		Type: MarshalMomentType(t.Type),
		Time: timestamppb.New(t.Time),
	}
}

func UnmarshalMoment(t *geckgov1.Moment) *Moment {
	if t == nil {
		return nil
	}

	return &Moment{
		Type: UnmarshalMomentType(t.Type),
		Time: t.Time.AsTime(),
	}
}

func MarshalMomentType(t MomentType) geckgov1.Moment_Type {
	return geckgov1.Moment_Type(t + 1)
}

func UnmarshalMomentType(t geckgov1.Moment_Type) MomentType {
	return MomentType(t - 1)
}

func MarshalDay(t *Day) *geckgov1.Day {
	if t == nil {
		return nil
	}

	moments := make([]*geckgov1.Moment, 0, len(t.Moments))
	for _, m := range t.Moments {
		moments = append(moments, MarshalMoment(&m))
	}

	return &geckgov1.Day{
		Date:    timestamppb.New(t.Date),
		Moments: moments,
	}
}

func UnmarshalDay(t *geckgov1.Day) *Day {
	if t == nil {
		return nil
	}

	moments := make([]Moment, 0, len(t.Moments))
	for _, m := range t.Moments {
		moments = append(moments, *UnmarshalMoment(m))
	}

	return &Day{
		Date:    t.Date.AsTime(),
		Moments: moments,
	}
}
