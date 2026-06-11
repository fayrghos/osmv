package utils

import rl "github.com/gen2brain/raylib-go/raylib"

func IsKeyPressedDouble(key int32) bool {
	return rl.IsKeyPressed(key) || rl.IsKeyPressedRepeat(key)
}
