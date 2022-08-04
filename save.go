package trimps

func (s *Save) ChangePlayerAction(pa Action) (ok bool) {
	if pa == s.PlayerAction {
		return true
	}
	if pa == Trapping && s.Player.Save.Buildings[Trap].Locked {
		return false
	}
	if pa == Lumbering && s.Player.Save.Resources[Wood].Locked {
		return false
	}
	if pa == Mining && s.Player.Save.Resources[Metal].Locked {
		return false
	}
	if pa == Researching && s.Player.Save.Resources[Science].Locked {
		return false
	}
	if pa == Farming && !s.Flag.Has(WoodFlag) {
		s.Player.PendingMessage = append(s.Player.PendingMessage, " 你需要有木头才能建造其他东西……\n使用 /玩家动作 砍伐 来弄一点木头。")
		s.Resources[Wood].Locked = false
		s.Flag.Set(WoodFlag)
	}
	if pa == Lumbering && !s.Flag.Has(TrapFlag) {
		s.Player.PendingMessage = append(s.Player.PendingMessage, " 或许你可以用陷阱抓到一些嘎嘣脆肉味的东西。\n使用 /建筑 来查看建筑列表， /建筑 陷阱 来开始陷阱项目，并使用 /玩家动作 建造 来开始建造你的陷阱。")
		s.Buildings[Trap].Locked = false
		s.Flag.Set(TrapFlag)
	}
	s.PlayerAction = pa
	return true
}
