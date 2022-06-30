package main

type Profile struct {
	ProfileName                     string  `toml:"profile_name"`
	DragFunction                    int     `toml:"drag_function"`
	Bc                              float64 `toml:"bc"`
	BulletWeight                    float64 `toml:"bullet_weight"`
	InitVelocity                    float64 `toml:"init_velocity"`
	SightHeight                     float64 `toml:"sight_height"`
	ZeroingDistance                 float64 `toml:"zeroing_distance"`
	ReticleOffsetX                  int     `toml:"reticle_offset_x"`
	ReticleOffsetY                  int     `toml:"reticle_offset_y"`
	MspAccBorderValue               int     `toml:"msp_acc_border_value"`
	AccBorderUpCrossCounterMin      int     `toml:"acc_border_up_cross_counter_min"`
	MspAccBorderUpCrossCounterMax   int     `toml:"msp_acc_border_up_cross_counter_max"`
	MspAccBorderDownCrossCounterMin int     `toml:"msp_acc_border_down_cross_counter_min"`
}

type TOMLSettings struct {
	SightHeightIn   bool    `toml:"sight_height_in"`
	VelocityFps     bool    `toml:"velocity_fps"`
	ZeroDistanceYds bool    `toml:"zero_distance_yds"`
	One             Profile `toml:"one"`
	Two             Profile `toml:"two"`
	Three           Profile `toml:"three"`
	Four            Profile `toml:"four"`
	Five            Profile `toml:"five"`
	Six             Profile `toml:"six"`
}

var XMLTemplate = `
<?xml version="1.0" encoding="utf-8"?>
<profiles_count>
<value_node_0>6</value_node_0>
</profiles_count>
<active_profile>
<value_node_0>0</value_node_0>
</active_profile>
<storage_drag_function>
<value_node_0>{{.One.DragFunction}}</value_node_0>
<value_node_1>{{.Two.DragFunction}}</value_node_1>
<value_node_2>{{.Three.DragFunction}}</value_node_2>
<value_node_3>{{.Four.DragFunction}}</value_node_3>
<value_node_4>{{.Five.DragFunction}}</value_node_4>
<value_node_5>{{.Six.DragFunction}}</value_node_5>
</storage_drag_function>
<storage_ballistic_coeff>
<value_node_0>{{.One.Bc}}</value_node_0>
<value_node_1>{{.Two.Bc}}</value_node_1>
<value_node_2>{{.Three.Bc}}</value_node_2>
<value_node_3>{{.Four.Bc}}</value_node_3>
<value_node_4>{{.Five.Bc}}</value_node_4>
<value_node_5>{{.Six.Bc}}</value_node_5>
</storage_ballistic_coeff>
<storage_bullet_weight>
<value_node_0>{{.One.BulletWeight}}</value_node_0>
<value_node_1>{{.Two.BulletWeight}}</value_node_1>
<value_node_2>{{.Three.BulletWeight}}</value_node_2>
<value_node_3>{{.Four.BulletWeight}}</value_node_3>\
<value_node_4>{{.Five.BulletWeight}}</value_node_4>
<value_node_5>{{.Six.BulletWeight}}</value_node_5>
</storage_bullet_weight>
<storage_init_velocity>
<value_node_0>{{.One.InitVelocity}}</value_node_0>
<value_node_1>{{.Two.InitVelocity}}</value_node_1>
<value_node_2>{{.Three.InitVelocity}}</value_node_2>
<value_node_3>{{.Four.InitVelocity}}</value_node_3>
<value_node_4>{{.Five.InitVelocity}}</value_node_4>
<value_node_5>{{.Six.InitVelocity}}</value_node_5>
</storage_init_velocity>
<storage_sight_height>
<value_node_0>{{.One.SightHeight}}</value_node_0>
<value_node_1>{{.Two.SightHeight}}</value_node_1>
<value_node_2>{{.Three.SightHeight}}</value_node_2>
<value_node_3>{{.Four.SightHeight}}</value_node_3>
<value_node_4>{{.Five.SightHeight}}</value_node_4>
<value_node_5>{{.Six.SightHeight}}</value_node_5>
</storage_sight_height>
<storage_zeroing_distance>
<value_node_0>{{.One.ZeroingDistance}}</value_node_0>
<value_node_1>{{.Two.ZeroingDistance}}</value_node_1>
<value_node_2>{{.Three.ZeroingDistance}}</value_node_2>
<value_node_3>{{.Four.ZeroingDistance}}</value_node_3>
<value_node_4>{{.Five.ZeroingDistance}}</value_node_4>
<value_node_5>{{.Six.ZeroingDistance}}</value_node_5>
</storage_zeroing_distance>
<storage_reticle_offset_x>
<value_node_0>{{.One.ReticleOffsetX}}</value_node_0>
<value_node_1>{{.Two.ReticleOffsetX}}</value_node_1>
<value_node_2>{{.Three.ReticleOffsetX}}</value_node_2>
<value_node_3>{{.Four.ReticleOffsetX}}</value_node_3>
<value_node_4>{{.Five.ReticleOffsetX}}</value_node_4>
<value_node_5>{{.Six.ReticleOffsetX}}</value_node_5>
</storage_reticle_offset_x>
<storage_reticle_offset_y>
<value_node_0>{{.One.ReticleOffsetY}}</value_node_0>
<value_node_1>{{.Two.ReticleOffsetY}}</value_node_1>
<value_node_2>{{.Three.ReticleOffsetY}}</value_node_2>
<value_node_3>{{.Four.ReticleOffsetY}}</value_node_3>
<value_node_4>{{.Five.ReticleOffsetY}}</value_node_4>
<value_node_5>{{.Six.ReticleOffsetY}}</value_node_5>
</storage_reticle_offset_y>
<storage_msp_acc_border_value>
<value_node_0>{{.One.MspAccBorderValue}}</value_node_0>
<value_node_1>{{.Two.MspAccBorderValue}}</value_node_1>
<value_node_2>{{.Three.MspAccBorderValue}}</value_node_2>
<value_node_3>{{.Four.MspAccBorderValue}}</value_node_3>
<value_node_4>{{.Five.MspAccBorderValue}}</value_node_4>
<value_node_5>{{.Six.MspAccBorderValue}}</value_node_5>
</storage_msp_acc_border_value>
<storage_msp_acc_border_up_cross_counter_min>
<value_node_0>{{.One.AccBorderUpCrossCounterMin}}</value_node_0>
<value_node_1>{{.Two.AccBorderUpCrossCounterMin}}</value_node_1>
<value_node_2>{{.Three.AccBorderUpCrossCounterMin}}</value_node_2>
<value_node_3>{{.Four.AccBorderUpCrossCounterMin}}</value_node_3>
<value_node_4>{{.Five.AccBorderUpCrossCounterMin}}</value_node_4>
<value_node_5>{{.Six.AccBorderUpCrossCounterMin}}</value_node_5>
</storage_msp_acc_border_up_cross_counter_min>
<storage_msp_acc_border_up_cross_counter_max>
<value_node_0>{{.One.MspAccBorderUpCrossCounterMax}}</value_node_0>
<value_node_1>{{.Two.MspAccBorderUpCrossCounterMax}}</value_node_1>
<value_node_2>{{.Three.MspAccBorderUpCrossCounterMax}}</value_node_2>
<value_node_3>{{.Four.MspAccBorderUpCrossCounterMax}}</value_node_3>
<value_node_4>{{.Five.MspAccBorderUpCrossCounterMax}}</value_node_4>
<value_node_5>{{.Six.MspAccBorderUpCrossCounterMax}}</value_node_5>
</storage_msp_acc_border_up_cross_counter_max>
<storage_msp_acc_border_down_cross_counter_min>
<value_node_0>{{.One.MspAccBorderDownCrossCounterMin}}</value_node_0>
<value_node_1>{{.Two.MspAccBorderDownCrossCounterMin}}</value_node_1>
<value_node_2>{{.Three.MspAccBorderDownCrossCounterMin}}</value_node_2>
<value_node_3>{{.Four.MspAccBorderDownCrossCounterMin}}</value_node_3>
<value_node_4>{{.Five.MspAccBorderDownCrossCounterMin}}</value_node_4>
<value_node_5>{{.Six.MspAccBorderDownCrossCounterMin}}</value_node_5>
</storage_msp_acc_border_down_cross_counter_min>
<storage_profile_name>
<value_node_0>{{.One.ProfileName}}</value_node_0>
<value_node_1>{{.Two.ProfileName}}</value_node_1>
<value_node_2>{{.Three.ProfileName}}</value_node_2>
<value_node_3>{{.Four.ProfileName}}</value_node_3>
<value_node_4>{{.Five.ProfileName}}</value_node_4>
<value_node_5>{{.Six.ProfileName}}</value_node_5>
</storage_profile_name>
`
