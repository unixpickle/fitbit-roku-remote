package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/model3d/model2d"
)

const (
	MarchingDelta = 0.001
	BarThickness  = 0.04
	BesselFrac    = 0.2
	IconSize      = 80
)

func main() {
	body := FinalObject(Rounded(
		model2d.NewRect(model2d.XY(-0.2, -0.5), model2d.XY(0.2, 0.5)),
		0.05,
	))
	bodyBorder := OutsetObject(body, 0.025)
	circle := FinalObject(&model2d.Circle{Center: model2d.Y(0.1), Radius: 0.15})
	arrows := FinalObject(Rounded(
		model2d.JoinedSolid{
			model2d.NewRect(model2d.XY(-0.1, 0.1-BarThickness), model2d.XY(0.1, 0.1+BarThickness)),
			model2d.NewRect(model2d.XY(-BarThickness, 0.0), model2d.XY(BarThickness, 0.2)),
		},
		0.02,
	))
	buttons := FinalObject(Rounded(
		model2d.JoinedSolid{
			model2d.NewRect(model2d.XY(-0.15, 0.35), model2d.XY(-0.02, 0.45)),
			model2d.NewRect(model2d.XY(0.02, 0.35), model2d.XY(0.15, 0.45)),
			model2d.NewRect(model2d.XY(-0.15, -0.45), model2d.XY(0.15, -0.35)),
			model2d.NewRect(model2d.XY(-0.15, -0.35+0.04), model2d.XY(0.15, -0.25+0.04)),
		},
		0.02,
	))

	sizeVec := body.Max().Sub(body.Min())
	size := math.Max(sizeVec.X, sizeVec.Y) * (1 + BesselFrac)
	mid := body.Max().Mid(body.Min())
	min := mid.AddScalar(-size / 2)
	max := mid.AddScalar(size / 2)

	rast := &model2d.Rasterizer{
		Bounds: model2d.NewRect(min, max),
		Scale:  IconSize / size,
	}
	bodyBorderImg := rast.RasterizeColliderSolid(bodyBorder)
	bodyImg := rast.RasterizeColliderSolid(body)
	circleImg := rast.RasterizeColliderSolid(circle)
	arrowsImg := rast.RasterizeColliderSolid(arrows)
	buttonsImg := rast.RasterizeColliderSolid(buttons)

	img := model2d.ColorizeOverlay(
		[]*image.Gray{bodyBorderImg, bodyImg, circleImg, arrowsImg, buttonsImg},
		[]color.Color{
			color.Gray{Y: 0x70},
			color.Gray{Y: 0x50},
			color.Gray{Y: 0x40},
			color.RGBA{R: 0x6f, G: 0x1a, B: 0xb1, A: 0xff},
			color.Gray{Y: 0x40},
		},
	)
	w, err := os.Create("../../resources/icon.png")
	essentials.Must(err)
	defer w.Close()
	essentials.Must(png.Encode(w, img))
}

func Rounded(s model2d.Solid, radius float64) model2d.Solid {
	m := model2d.MarchingSquaresSearch(s, MarchingDelta, 8)
	s1 := model2d.NewColliderSolidInset(model2d.MeshToCollider(m), radius)
	m2 := model2d.MarchingSquaresSearch(s1, MarchingDelta, 8)
	s2 := model2d.NewColliderSolidInset(model2d.MeshToCollider(m2), -radius)
	return s2
}

func FinalObject(s model2d.Solid) model2d.Collider {
	m := model2d.MarchingSquaresSearch(s, MarchingDelta, 8)
	return model2d.MeshToCollider(m.MapCoords(model2d.XY(1.0, -1.0).Mul).Rotate(math.Pi / 6))
}

func OutsetObject(c model2d.Collider, outset float64) model2d.Collider {
	m := model2d.NewColliderSolidInset(c, -outset)
	return model2d.MeshToCollider(model2d.MarchingSquaresSearch(m, MarchingDelta, 8))
}
