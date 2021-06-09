package geos

import (
	"github.com/flywave/go-geom"
)

func ConvertGeomToGeos(g geom.Geometry) *Geometry {
	switch t := g.(type) {
	case geom.Point:
		return CreatePoint(t.X(), t.Y())
	case geom.Point3:
		return CreatePointZ(t.X(), t.Y(), t.Z())
	case geom.MultiPoint:
	case geom.MultiPoint3:
		geos := make([]*Geometry, len(t.Points()))
		for i := range t.Points() {
			geos[i] = ConvertGeomToGeos(t.Points()[i])
		}
		return CreateMultiGeometry(geos, MULTIPOINT)
	case geom.LineString:
		coords := make([]Coord, len(t.Subpoints()))
		l := t.Subpoints()
		for i := range l {
			coords[i] = Coord{l[i].X(), l[i].Y()}
		}
		return CreateLineString(coords)
	case geom.LineString3:
		coords := make([]CoordZ, len(t.Subpoints()))
		l := t.Subpoints()
		for i := range l {
			coords[i] = CoordZ{l[i].X(), l[i].Y(), l[i].Z()}
		}
		return CreateLineStringZ(coords)
	case geom.MultiLine:
	case geom.MultiLine3:
		geos := make([]*Geometry, len(t.Lines()))
		for i := range t.Lines() {
			geos[i] = ConvertGeomToGeos(t.Lines()[i])
		}
		return CreateMultiGeometry(geos, MULTILINESTRING)
	case geom.Polygon:
		ls := t.Sublines()
		lps := ls[0].Subpoints()
		shell := make([]Coord, len(lps))
		for i := range shell {
			shell[i] = Coord{lps[i].X(), lps[i].Y()}
		}
		subcoords := make([][]Coord, 0, len(ls[1:]))
		for l := range ls[1:] {
			lps = ls[l].Subpoints()
			coords := make([]Coord, len(lps))
			for i := range coords {
				coords[i] = Coord{lps[i].X(), lps[i].Y()}
			}
			subcoords = append(subcoords, coords)
		}
		return CreatePolygon(shell, subcoords...)
	case geom.Polygon3:
		ls := t.Sublines()
		lps := ls[0].Subpoints()
		shell := make([]CoordZ, len(lps))
		for i := range shell {
			shell[i] = CoordZ{lps[i].X(), lps[i].Y(), lps[i].Z()}
		}
		subcoords := make([][]CoordZ, 0, len(ls[1:]))
		for l := range ls[1:] {
			lps = ls[l].Subpoints()
			coords := make([]CoordZ, len(lps))
			for i := range coords {
				coords[i] = CoordZ{lps[i].X(), lps[i].Y(), lps[i].Z()}
			}
			subcoords = append(subcoords, coords)
		}
		return CreatePolygonZ(shell, subcoords...)
	case geom.MultiPolygon:
	case geom.MultiPolygon3:
		geos := make([]*Geometry, len(t.Polygons()))
		for i := range t.Polygons() {
			geos[i] = ConvertGeomToGeos(t.Polygons()[i])
		}
		return CreateMultiGeometry(geos, MULTIPOLYGON)
	}
	return nil
}

type GeosPoint struct {
	geom.Point
	g *Geometry
}

func (p *GeosPoint) GetType() string {
	return string(geom.GeometryPoint)
}

func (p *GeosPoint) X() float64 {
	x, _ := p.g.GetXY()
	return x
}

func (p *GeosPoint) Y() float64 {
	_, y := p.g.GetXY()
	return y
}

func (p *GeosPoint) Data() []float64 {
	x, y := p.g.GetXY()
	return []float64{x, y}
}

type GeosPoint3 struct {
	geom.Point3
	g *Geometry
}

func (p *GeosPoint3) GetType() string {
	return string(geom.GeometryPoint)
}

func (p *GeosPoint3) X() float64 {
	x, _ := p.g.GetXY()
	return x
}

func (p *GeosPoint3) Y() float64 {
	_, y := p.g.GetXY()
	return y
}

func (p *GeosPoint3) Z() float64 {
	return p.g.GetZ()
}

func (p *GeosPoint3) Data() []float64 {
	x, y := p.g.GetXY()
	return []float64{x, y, p.g.GetZ()}
}

type GeosMultiPoint struct {
	geom.MultiPoint
	g *Geometry
}

func (p *GeosMultiPoint) GetType() string {
	return string(geom.GeometryMultiPoint)
}

func (p *GeosMultiPoint) Points() []geom.Point {
	ret := make([]geom.Point, p.g.GetNumGeometries())
	for i := range ret {
		ret[i] = &GeosPoint{g: p.g.GetGeometryN(i)}
	}
	return ret
}

func (p *GeosMultiPoint) Data() [][]float64 {
	ret := make([][]float64, p.g.GetNumGeometries())
	for i := range ret {
		r := &GeosPoint{g: p.g.GetGeometryN(i)}
		ret[i] = r.Data()
	}
	return ret
}

type GeosMultiPoint3 struct {
	geom.MultiPoint3
	g *Geometry
}

func (p *GeosMultiPoint3) GetType() string {
	return string(geom.GeometryMultiPoint)
}

func (p *GeosMultiPoint3) Points() []geom.Point3 {
	ret := make([]geom.Point3, p.g.GetNumGeometries())
	for i := range ret {
		ret[i] = &GeosPoint3{g: p.g.GetGeometryN(i)}
	}
	return ret
}

func (p *GeosMultiPoint3) Data() [][]float64 {
	ret := make([][]float64, p.g.GetNumGeometries())
	for i := range ret {
		r := &GeosPoint3{g: p.g.GetGeometryN(i)}
		ret[i] = r.Data()
	}
	return ret
}

type GeosCoordPoint struct {
	geom.Point
	c Coord
}

func (p *GeosCoordPoint) GetType() string {
	return string(geom.GeometryPoint)
}

func (p *GeosCoordPoint) X() float64 {
	return p.c.X
}

func (p *GeosCoordPoint) Y() float64 {
	return p.c.Y
}

func (p *GeosCoordPoint) Data() []float64 {
	return []float64{p.c.X, p.c.Y}
}

type GeosCoordPoint3 struct {
	geom.Point3
	c CoordZ
}

func (p *GeosCoordPoint3) GetType() string {
	return string(geom.GeometryPoint)
}

func (p *GeosCoordPoint3) X() float64 {
	return p.c.X
}

func (p *GeosCoordPoint3) Y() float64 {
	return p.c.Y
}

func (p *GeosCoordPoint3) Z() float64 {
	return p.c.Z
}

func (p *GeosCoordPoint3) Data() []float64 {
	return []float64{p.c.X, p.c.Y, p.c.Z}
}

type GeosLineString struct {
	geom.LineString
	g *Geometry
}

func (p *GeosLineString) GetType() string {
	return string(geom.GeometryLineString)
}

func (p *GeosLineString) Subpoints() []geom.Point {
	coords := p.g.GetCoords()
	ret := make([]geom.Point, len(coords))
	for i := range ret {
		ret[i] = &GeosCoordPoint{c: coords[i]}
	}
	return ret
}

func (p *GeosLineString) Data() [][]float64 {
	coords := p.g.GetCoords()
	ret := make([][]float64, len(coords))
	for i := range ret {
		r := &GeosCoordPoint{c: coords[i]}
		ret[i] = r.Data()
	}
	return ret
}

type GeosLineString3 struct {
	geom.LineString3
	g *Geometry
}

func (p *GeosLineString3) GetType() string {
	return string(geom.GeometryLineString)
}

func (p *GeosLineString3) Subpoints() []geom.Point3 {
	coords := p.g.GetCoordZs()
	ret := make([]geom.Point3, len(coords))
	for i := range ret {
		ret[i] = &GeosCoordPoint3{c: coords[i]}
	}
	return ret
}

func (p *GeosLineString3) Data() [][]float64 {
	coords := p.g.GetCoordZs()
	ret := make([][]float64, len(coords))
	for i := range ret {
		r := &GeosCoordPoint3{c: coords[i]}
		ret[i] = r.Data()
	}
	return ret
}

type GeosMultiLine struct {
	geom.MultiLine
	g *Geometry
}

func (p *GeosMultiLine) GetType() string {
	return string(geom.GeometryMultiLineString)
}

func (p *GeosMultiLine) Lines() []geom.LineString {
	ret := make([]geom.LineString, p.g.GetNumGeometries())
	for i := range ret {
		ret[i] = &GeosLineString{g: p.g.GetGeometryN(i)}
	}
	return ret
}

func (p *GeosMultiLine) Data() [][][]float64 {
	ret := make([][][]float64, p.g.GetNumGeometries())
	for i := range ret {
		r := &GeosLineString{g: p.g.GetGeometryN(i)}
		ret[i] = r.Data()
	}
	return ret
}

type GeosMultiLine3 struct {
	geom.MultiLine3
	g *Geometry
}

func (p *GeosMultiLine3) GetType() string {
	return string(geom.GeometryMultiLineString)
}

func (p *GeosMultiLine3) Lines() []geom.LineString3 {
	ret := make([]geom.LineString3, p.g.GetNumGeometries())
	for i := range ret {
		ret[i] = &GeosLineString3{g: p.g.GetGeometryN(i)}
	}
	return ret
}

func (p *GeosMultiLine3) Data() [][][]float64 {
	ret := make([][][]float64, p.g.GetNumGeometries())
	for i := range ret {
		r := &GeosLineString3{g: p.g.GetGeometryN(i)}
		ret[i] = r.Data()
	}
	return ret
}

type GeosPolygon struct {
	geom.Polygon
	g *Geometry
}

func (p *GeosPolygon) GetType() string {
	return string(geom.GeometryPolygon)
}

func (p *GeosPolygon) Sublines() []geom.LineString {
	ret := make([]geom.LineString, p.g.GetNumInteriorRings()+1)
	ret[0] = &GeosLineString{g: p.g.GetExteriorRing()}
	rings := ret[1:]
	for i := range rings {
		rings[i] = &GeosLineString{g: p.g.GetInteriorRingN(i)}
	}
	return rings
}

func (p *GeosPolygon) Data() [][][]float64 {
	ret := make([][][]float64, p.g.GetNumInteriorRings()+1)
	r := &GeosLineString{g: p.g.GetExteriorRing()}
	ret[0] = r.Data()
	rings := ret[1:]
	for i := range rings {
		r = &GeosLineString{g: p.g.GetInteriorRingN(i)}
		rings[i] = r.Data()
	}
	return ret
}

type GeosPolygon3 struct {
	geom.Polygon3
	g *Geometry
}

func (p *GeosPolygon3) GetType() string {
	return string(geom.GeometryPolygon)
}

func (p *GeosPolygon3) Sublines() []geom.LineString3 {
	ret := make([]geom.LineString3, p.g.GetNumInteriorRings()+1)
	ret[0] = &GeosLineString3{g: p.g.GetExteriorRing()}
	rings := ret[1:]
	for i := range rings {
		rings[i] = &GeosLineString3{g: p.g.GetInteriorRingN(i)}
	}
	return rings
}

func (p *GeosPolygon3) Data() [][][]float64 {
	ret := make([][][]float64, p.g.GetNumInteriorRings()+1)
	r := &GeosLineString3{g: p.g.GetExteriorRing()}
	ret[0] = r.Data()
	rings := ret[1:]
	for i := range rings {
		r = &GeosLineString3{g: p.g.GetInteriorRingN(i)}
		rings[i] = r.Data()
	}
	return ret
}

type GeosMultiPolygon struct {
	geom.MultiPolygon
	g *Geometry
}

func (p *GeosMultiPolygon) GetType() string {
	return string(geom.GeometryMultiPolygon)
}

func (p *GeosMultiPolygon) Polygons() []geom.Polygon {
	ret := make([]geom.Polygon, p.g.GetNumGeometries())
	for i := range ret {
		ret[i] = &GeosPolygon{g: p.g.GetGeometryN(i)}
	}
	return ret
}

func (p *GeosMultiPolygon) Data() [][][][]float64 {
	ret := make([][][][]float64, p.g.GetNumGeometries())
	for i := range ret {
		r := &GeosPolygon{g: p.g.GetGeometryN(i)}
		ret[i] = r.Data()
	}
	return ret
}

type GeosMultiPolygon3 struct {
	geom.MultiPolygon3
	g *Geometry
}

func (p *GeosMultiPolygon3) GetType() string {
	return string(geom.GeometryMultiPolygon)
}

func (p *GeosMultiPolygon3) Polygons() []geom.Polygon3 {
	ret := make([]geom.Polygon3, p.g.GetNumGeometries())
	for i := range ret {
		ret[i] = &GeosPolygon3{g: p.g.GetGeometryN(i)}
	}
	return ret
}

func (p *GeosMultiPolygon3) Data() [][][][]float64 {
	ret := make([][][][]float64, p.g.GetNumGeometries())
	for i := range ret {
		r := &GeosPolygon3{g: p.g.GetGeometryN(i)}
		ret[i] = r.Data()
	}
	return ret
}

func ConvertGeosToGeom(g *Geometry) geom.Geometry {
	switch g.GetType() {
	case POINT:
		if g.GetDimensions() == 2 {
			return &GeosPoint{g: g}
		} else if g.GetDimensions() == 3 {
			return &GeosPoint3{g: g}
		}
	case LINEARRING:
	case LINESTRING:
		if g.GetDimensions() == 2 {
			return &GeosLineString{g: g}
		} else if g.GetDimensions() == 3 {
			return &GeosLineString3{g: g}
		}
	case POLYGON:
		if g.GetDimensions() == 2 {
			return &GeosPolygon{g: g}
		} else if g.GetDimensions() == 3 {
			return &GeosPolygon3{g: g}
		}
	case MULTIPOINT:
		if g.GetDimensions() == 2 {
			return &GeosMultiPoint{g: g}
		} else if g.GetDimensions() == 3 {
			return &GeosMultiPoint3{g: g}
		}
	case MULTILINESTRING:
		if g.GetDimensions() == 2 {
			return &GeosMultiLine{g: g}
		} else if g.GetDimensions() == 3 {
			return &GeosMultiLine3{g: g}
		}
	case MULTIPOLYGON:
		if g.GetDimensions() == 2 {
			return &GeosMultiPolygon{g: g}
		} else if g.GetDimensions() == 3 {
			return &GeosMultiPolygon3{g: g}
		}
	case GEOMETRYCOLLECTION:
		ret := make(geom.Collection, g.GetNumGeometries())
		for i := range ret {
			ret[i] = ConvertGeosToGeom(g.GetGeometryN(i))
		}
		return ret
	default:
		return nil
	}
	return nil
}
