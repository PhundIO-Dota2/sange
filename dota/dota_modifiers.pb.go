// Code generated by protoc-gen-go.
// source: dota_modifiers.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import google_protobuf "github.com/dotabuff/sange/dota/google/protobuf/descriptor.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type DOTA_MODIFIER_ENTRY_TYPE int32

const (
	DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_ACTIVE  DOTA_MODIFIER_ENTRY_TYPE = 1
	DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_REMOVED DOTA_MODIFIER_ENTRY_TYPE = 2
)

var DOTA_MODIFIER_ENTRY_TYPE_name = map[int32]string{
	1: "DOTA_MODIFIER_ENTRY_TYPE_ACTIVE",
	2: "DOTA_MODIFIER_ENTRY_TYPE_REMOVED",
}
var DOTA_MODIFIER_ENTRY_TYPE_value = map[string]int32{
	"DOTA_MODIFIER_ENTRY_TYPE_ACTIVE":  1,
	"DOTA_MODIFIER_ENTRY_TYPE_REMOVED": 2,
}

func (x DOTA_MODIFIER_ENTRY_TYPE) Enum() *DOTA_MODIFIER_ENTRY_TYPE {
	p := new(DOTA_MODIFIER_ENTRY_TYPE)
	*p = x
	return p
}
func (x DOTA_MODIFIER_ENTRY_TYPE) String() string {
	return proto.EnumName(DOTA_MODIFIER_ENTRY_TYPE_name, int32(x))
}
func (x *DOTA_MODIFIER_ENTRY_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DOTA_MODIFIER_ENTRY_TYPE_value, data, "DOTA_MODIFIER_ENTRY_TYPE")
	if err != nil {
		return err
	}
	*x = DOTA_MODIFIER_ENTRY_TYPE(value)
	return nil
}

type CDOTAModifierBuffTableEntry struct {
	EntryType           *DOTA_MODIFIER_ENTRY_TYPE `protobuf:"varint,1,req,name=entry_type,enum=dota.DOTA_MODIFIER_ENTRY_TYPE,def=1" json:"entry_type,omitempty"`
	Parent              *int32                    `protobuf:"varint,2,req,name=parent" json:"parent,omitempty"`
	Index               *int32                    `protobuf:"varint,3,req,name=index" json:"index,omitempty"`
	SerialNum           *int32                    `protobuf:"varint,4,req,name=serial_num" json:"serial_num,omitempty"`
	ModifierClass       *int32                    `protobuf:"varint,5,opt,name=modifier_class" json:"modifier_class,omitempty"`
	AbilityLevel        *int32                    `protobuf:"varint,6,opt,name=ability_level" json:"ability_level,omitempty"`
	StackCount          *int32                    `protobuf:"varint,7,opt,name=stack_count" json:"stack_count,omitempty"`
	CreationTime        *float32                  `protobuf:"fixed32,8,opt,name=creation_time" json:"creation_time,omitempty"`
	Duration            *float32                  `protobuf:"fixed32,9,opt,name=duration,def=-1" json:"duration,omitempty"`
	Caster              *int32                    `protobuf:"varint,10,opt,name=caster" json:"caster,omitempty"`
	Ability             *int32                    `protobuf:"varint,11,opt,name=ability" json:"ability,omitempty"`
	Armor               *int32                    `protobuf:"varint,12,opt,name=armor" json:"armor,omitempty"`
	FadeTime            *float32                  `protobuf:"fixed32,13,opt,name=fade_time" json:"fade_time,omitempty"`
	Subtle              *bool                     `protobuf:"varint,14,opt,name=subtle" json:"subtle,omitempty"`
	ChannelTime         *float32                  `protobuf:"fixed32,15,opt,name=channel_time" json:"channel_time,omitempty"`
	VStart              *CMsgVector               `protobuf:"bytes,16,opt,name=v_start" json:"v_start,omitempty"`
	VEnd                *CMsgVector               `protobuf:"bytes,17,opt,name=v_end" json:"v_end,omitempty"`
	PortalLoopAppear    *string                   `protobuf:"bytes,18,opt,name=portal_loop_appear" json:"portal_loop_appear,omitempty"`
	PortalLoopDisappear *string                   `protobuf:"bytes,19,opt,name=portal_loop_disappear" json:"portal_loop_disappear,omitempty"`
	HeroLoopAppear      *string                   `protobuf:"bytes,20,opt,name=hero_loop_appear" json:"hero_loop_appear,omitempty"`
	HeroLoopDisappear   *string                   `protobuf:"bytes,21,opt,name=hero_loop_disappear" json:"hero_loop_disappear,omitempty"`
	MovementSpeed       *int32                    `protobuf:"varint,22,opt,name=movement_speed" json:"movement_speed,omitempty"`
	Aura                *bool                     `protobuf:"varint,23,opt,name=aura" json:"aura,omitempty"`
	Activity            *int32                    `protobuf:"varint,24,opt,name=activity" json:"activity,omitempty"`
	Damage              *int32                    `protobuf:"varint,25,opt,name=damage" json:"damage,omitempty"`
	Range               *int32                    `protobuf:"varint,26,opt,name=range" json:"range,omitempty"`
	DdModifierIndex     *int32                    `protobuf:"varint,27,opt,name=dd_modifier_index" json:"dd_modifier_index,omitempty"`
	DdAbilityIndex      *int32                    `protobuf:"varint,28,opt,name=dd_ability_index" json:"dd_ability_index,omitempty"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *CDOTAModifierBuffTableEntry) Reset()         { *m = CDOTAModifierBuffTableEntry{} }
func (m *CDOTAModifierBuffTableEntry) String() string { return proto.CompactTextString(m) }
func (*CDOTAModifierBuffTableEntry) ProtoMessage()    {}

const Default_CDOTAModifierBuffTableEntry_EntryType DOTA_MODIFIER_ENTRY_TYPE = DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_ACTIVE
const Default_CDOTAModifierBuffTableEntry_Duration float32 = -1

func (m *CDOTAModifierBuffTableEntry) GetEntryType() DOTA_MODIFIER_ENTRY_TYPE {
	if m != nil && m.EntryType != nil {
		return *m.EntryType
	}
	return Default_CDOTAModifierBuffTableEntry_EntryType
}

func (m *CDOTAModifierBuffTableEntry) GetParent() int32 {
	if m != nil && m.Parent != nil {
		return *m.Parent
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetSerialNum() int32 {
	if m != nil && m.SerialNum != nil {
		return *m.SerialNum
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetModifierClass() int32 {
	if m != nil && m.ModifierClass != nil {
		return *m.ModifierClass
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetAbilityLevel() int32 {
	if m != nil && m.AbilityLevel != nil {
		return *m.AbilityLevel
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetStackCount() int32 {
	if m != nil && m.StackCount != nil {
		return *m.StackCount
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetCreationTime() float32 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDuration() float32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return Default_CDOTAModifierBuffTableEntry_Duration
}

func (m *CDOTAModifierBuffTableEntry) GetCaster() int32 {
	if m != nil && m.Caster != nil {
		return *m.Caster
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetAbility() int32 {
	if m != nil && m.Ability != nil {
		return *m.Ability
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetArmor() int32 {
	if m != nil && m.Armor != nil {
		return *m.Armor
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetFadeTime() float32 {
	if m != nil && m.FadeTime != nil {
		return *m.FadeTime
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetSubtle() bool {
	if m != nil && m.Subtle != nil {
		return *m.Subtle
	}
	return false
}

func (m *CDOTAModifierBuffTableEntry) GetChannelTime() float32 {
	if m != nil && m.ChannelTime != nil {
		return *m.ChannelTime
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetVStart() *CMsgVector {
	if m != nil {
		return m.VStart
	}
	return nil
}

func (m *CDOTAModifierBuffTableEntry) GetVEnd() *CMsgVector {
	if m != nil {
		return m.VEnd
	}
	return nil
}

func (m *CDOTAModifierBuffTableEntry) GetPortalLoopAppear() string {
	if m != nil && m.PortalLoopAppear != nil {
		return *m.PortalLoopAppear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetPortalLoopDisappear() string {
	if m != nil && m.PortalLoopDisappear != nil {
		return *m.PortalLoopDisappear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetHeroLoopAppear() string {
	if m != nil && m.HeroLoopAppear != nil {
		return *m.HeroLoopAppear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetHeroLoopDisappear() string {
	if m != nil && m.HeroLoopDisappear != nil {
		return *m.HeroLoopDisappear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetMovementSpeed() int32 {
	if m != nil && m.MovementSpeed != nil {
		return *m.MovementSpeed
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetAura() bool {
	if m != nil && m.Aura != nil {
		return *m.Aura
	}
	return false
}

func (m *CDOTAModifierBuffTableEntry) GetActivity() int32 {
	if m != nil && m.Activity != nil {
		return *m.Activity
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDamage() int32 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetRange() int32 {
	if m != nil && m.Range != nil {
		return *m.Range
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDdModifierIndex() int32 {
	if m != nil && m.DdModifierIndex != nil {
		return *m.DdModifierIndex
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDdAbilityIndex() int32 {
	if m != nil && m.DdAbilityIndex != nil {
		return *m.DdAbilityIndex
	}
	return 0
}

func init() {
	proto.RegisterEnum("dota.DOTA_MODIFIER_ENTRY_TYPE", DOTA_MODIFIER_ENTRY_TYPE_name, DOTA_MODIFIER_ENTRY_TYPE_value)
}
