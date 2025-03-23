package characters

import (
	"context"
	"strconv"
	"strings"
	"warhoop/app/model/char"
	"warhoop/app/utils"
)

func (svc *CharService) GetByID(ctx context.Context, id int) (*char.Characters, error) {
	result, err := svc.store.CharRepo.GetByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}

	return result.ToWeb(), nil
}

func (svc *CharService) GetByName(ctx context.Context, name string) (*char.Characters, error) {
	result, err := svc.store.CharRepo.GetByName(ctx, name)
	if err != nil {
		return nil, utils.ErrDataBase
	}

	inv := result.Inventory[:0]
	for _, equip := range result.Inventory {
		if equip.Slot >= 0 && equip.Slot <= 18 && equip.Bag == 0 {
			equip.Slot += 1
			inv = append(inv, equip)
		}
	}
	result.Inventory = inv

	for i := range result.Inventory {
		equip := &result.Inventory[i]
		if equip.ItemInstance != nil && equip.ItemInstance.ItemDBC != nil &&
			equip.ItemInstance.ItemDBC.ItemDisplayInfoID != 0 {

			switch equip.Slot {
			case 16:
				equip.Slot = 21
			case 17:
				equip.Slot = 22
			case 15:
				equip.Slot = 16
			case 5:
				if equip.ItemInstance.ItemDBC.InventoryType == 20 {
					equip.Slot = 20
				}
			}
		}
	}

	enchMap := make(map[int32][]int)

	for i := range result.Inventory {
		equip := &result.Inventory[i]

		if equip.ItemInstance != nil {
			enchValues := strings.Fields(equip.ItemInstance.Enchantments)
			if len(enchValues) > 0 {
				for _, field := range enchValues {
					id, err := strconv.ParseInt(field, 10, 32)
					if err == nil && id != 0 {
						enchMap[int32(id)] = append(enchMap[int32(id)], i)
					}
				}

				equip.ItemInstance.Enchantments = enchValues[0]
			} else {
				equip.ItemInstance.Enchantments = " "
			}
		}
	}

	var enchFullList []int32
	for enchID := range enchMap {
		if enchID != 3729 && enchID != 3717 {
			enchFullList = append(enchFullList, enchID)
		}
	}

	gemMap, err := svc.store.NexusRepo.GetEnchantDBCByIDs(ctx, enchFullList)
	if err != nil {
		return nil, utils.ErrDataBase
	}

	for enchID, itemIndexes := range enchMap {
		for _, itemIndex := range itemIndexes {
			equip := &result.Inventory[itemIndex]

			if enchID == 3729 || enchID == 3717 {
				equip.ItemInstance.Socket = strconv.Itoa(int(enchID))
				continue
			}

			if srcItemID, found := gemMap[enchID]; found {
				if equip.ItemInstance.Gems == "" {
					equip.ItemInstance.Gems = strconv.Itoa(int(srcItemID))
				} else {
					equip.ItemInstance.Gems += ":" + strconv.Itoa(int(srcItemID))
				}
			}
		}
	}

	return result.ToWeb(), nil
}

func (svc *CharService) GetTop10Kill(ctx context.Context) ([]map[string]interface{}, error) {
	result, err := svc.store.CharRepo.GetTop10Kill(ctx)
	if err != nil {
		return nil, utils.ErrDataBase
	}

	transformed := make([]map[string]interface{}, 0, len(*result))
	for _, char := range *result {
		transformed = append(transformed, map[string]interface{}{
			"name":       char.Name,
			"race":       char.Race,
			"class":      char.Class,
			"guid":       char.Guid,
			"totalkills": char.TotalKills,
		})
	}
	return transformed, nil
}

func (svc *CharService) GetOnlineCount(ctx context.Context) (int, error) {
	count, err := svc.store.CharRepo.GetOnlineCount(ctx)
	if err != nil {
		return 0, utils.ErrDataBase
	}
	return count, err
}

func (svc *CharService) GetOnlineSlice(ctx context.Context) ([]map[string]interface{}, error) {

	result, err := svc.store.CharRepo.GetOnlineSlice(ctx)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	transformed := make([]map[string]interface{}, 0, len(*result))
	for _, char := range *result {
		transformed = append(transformed, map[string]interface{}{
			"name":   char.Name,
			"level":  char.Level,
			"race":   char.Race,
			"class":  char.Class,
			"gender": char.Gender,
			"map":    char.Maps.ToWeb(),
			"zone":   char.Zones.ToWeb(),
		})
	}
	return transformed, err
}
