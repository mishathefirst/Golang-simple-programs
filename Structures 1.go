type CoolStruct struct {
    On bool
    Ammo int
    Power int
    
}

func (c *CoolStruct) Shoot() bool {
    if c.On == false {
        return false
    } else if c.Ammo > 0 {
        c.Ammo--
        return true
    } else {
        return false
    }
}

func (c *CoolStruct) RideBike()  bool {
    if c.On == false {
        return false
    } else if c.Power > 0 {
        c.Power--
        return true
    } else {
        return false
    }
}

