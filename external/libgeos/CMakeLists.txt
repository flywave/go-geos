IF(UNIX)
  SET(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wno-float-conversion -Wno-shadow -Wno-unused-function -Wno-format-extra-args -Wno-unknown-pragmas -Wno-tautological-compare -Wno-attributes -Wno-unused-const-variable -Wno-sign-compare -Wno-deprecated-declarations")
ENDIF()

IF(APPLE)
 SET(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=gnu++14")
ENDIF()	

if(UNIX AND NOT APPLE)
  SET(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fPIC ")
  SET(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fPIC ")
ENDIF()

INCLUDE_DIRECTORIES(BEFORE ${CMAKE_CURRENT_SOURCE_DIR}/geos/include)
INCLUDE_DIRECTORIES(BEFORE ${CMAKE_CURRENT_SOURCE_DIR}/geos/capi)
INCLUDE_DIRECTORIES(BEFORE ${CMAKE_CURRENT_SOURCE_DIR}/arch/common)

IF (WIN32 AND NOT MINGW)
  INCLUDE_DIRECTORIES(BEFORE ${CMAKE_CURRENT_SOURCE_DIR}/arch/win)
  set(CONFIG_SRC ${CMAKE_CURRENT_SOURCE_DIR}/arch/win/geos/platform.h)
ELSE()
  INCLUDE_DIRECTORIES(BEFORE ${CMAKE_CURRENT_SOURCE_DIR}/arch/unix)
  set(CONFIG_SRC ${CMAKE_CURRENT_SOURCE_DIR}/arch/unix/geos/platform.h)
ENDIF()


SET(SRC_GEOS
  ${CMAKE_CURRENT_SOURCE_DIR}/geos/capi/geos_c.cpp
  ${CMAKE_CURRENT_SOURCE_DIR}/geos/capi/geos_ts_c.cpp
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/inlines.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/Angle.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/BoundaryNodeRule.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/Centroid.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/CentroidArea.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/CentroidLine.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/CentroidPoint.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/CGAlgorithms.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/ConvexHull.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/HCoordinate.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/InteriorPointArea.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/InteriorPointLine.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/InteriorPointPoint.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/LineIntersector.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/MCPointInRing.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/MinimumDiameter.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/NotRepresentableException.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/PointLocator.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/RayCrossingCounter.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/RobustDeterminant.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/SimplePointInRing.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/algorithm/SIRtreePointInRing.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/Coordinate.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/CoordinateSequence.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/CoordinateSequenceFactory.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/CoordinateArraySequence.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/CoordinateArraySequenceFactory.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/Dimension.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/Envelope.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/Geometry.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/GeometryCollection.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/GeometryComponentFilter.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/GeometryFactory.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/IntersectionMatrix.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/LinearRing.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/LineSegment.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/LineString.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/Location.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/MultiLineString.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/MultiPoint.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/MultiPolygon.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/Point.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/Polygon.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/PrecisionModel.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/Triangle.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/AbstractPreparedPolygonContains.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/BasicPreparedGeometry.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/PreparedGeometry.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/PreparedGeometryFactory.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/PreparedLineString.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/PreparedLineStringIntersects.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/PreparedPoint.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/PreparedPolygonContains.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/PreparedPolygonContainsProperly.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/PreparedPolygonCovers.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/PreparedPolygon.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/PreparedPolygonIntersects.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/prep/PreparedPolygonPredicate.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/util/ComponentCoordinateExtracter.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/util/CoordinateOperation.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/util/GeometryEditor.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/util/GeometryTransformer.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/util/ShortCircuitedGeometryVisitor.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/util/SineStarFactory.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/util/GeometryCombiner.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/util/LinearComponentExtracter.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/util/PointExtracter.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geom/util/PolygonExtracter.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/Depth.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/DirectedEdge.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/DirectedEdgeStar.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/Edge.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/EdgeEnd.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/EdgeEndStar.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/EdgeIntersectionList.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/EdgeNodingValidator.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/EdgeList.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/EdgeRing.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/GeometryGraph.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/GraphComponent.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/Label.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/Node.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/NodeFactory.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/NodeMap.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/PlanarGraph.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/Position.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/Quadrant.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/TopologyLocation.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/index/MonotoneChainEdge.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/index/MonotoneChainIndexer.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/index/SegmentIntersector.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/index/SimpleEdgeSetIntersector.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/index/SimpleMCSweepLineIntersector.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/index/SimpleSweepLineIntersector.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/index/SweepLineEvent.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/geomgraph/index/SweepLineSegment.cpp 

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/io/ParseException.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/io/StringTokenizer.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/io/ByteOrderDataInStream.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/io/ByteOrderValues.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/io/WKTReader.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/io/WKTWriter.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/io/WKBReader.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/io/WKBWriter.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/io/Writer.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/io/Unload.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/io/CLocalizer.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/linearref/ExtractLineByLocation.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/linearref/LengthIndexedLine.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/linearref/LengthIndexOfPoint.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/linearref/LengthLocationMap.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/linearref/LinearIterator.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/linearref/LinearGeometryBuilder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/linearref/LinearLocation.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/linearref/LocationIndexOfLine.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/linearref/LocationIndexOfPoint.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/BasicSegmentString.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/FastNodingValidator.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/FastSegmentSetIntersectionFinder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/GeometryNoder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/IntersectionAdder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/IntersectionFinderAdder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/IteratedNoder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/MCIndexNoder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/MCIndexSegmentSetMutualIntersector.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/NodedSegmentString.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/NodingValidator.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/Octant.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/OrientedCoordinateArray.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/ScaledNoder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/SegmentIntersectionDetector.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/SegmentNode.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/SegmentNodeList.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/SegmentString.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/SegmentStringUtil.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/SimpleNoder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/SingleInteriorIntersectionFinder.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/snapround/HotPixel.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/snapround/MCIndexPointSnapper.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/snapround/MCIndexSnapRounder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/noding/snapround/SimpleSnapRounder.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/buffer/BufferBuilder.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/buffer/BufferInputLineSimplifier.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/buffer/BufferOp.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/buffer/BufferParameters.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/buffer/BufferSubgraph.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/buffer/OffsetCurveBuilder.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/buffer/OffsetCurveSetBuilder.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/buffer/OffsetSegmentGenerator.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/buffer/RightmostEdgeFinder.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/buffer/SubgraphDepthLocater.cpp 

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/distance/ConnectedElementLocationFilter.cpp 
    ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/distance/ConnectedElementPointFilter.cpp 
    ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/distance/DistanceOp.cpp 
    ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/distance/IndexedFacetDistance.cpp 
    ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/distance/FacetSequence.cpp 
    ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/distance/FacetSequenceTreeBuilder.cpp 
    ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/distance/GeometryLocation.cpp 

    ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/intersection/Rectangle.cpp 
    ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/intersection/RectangleIntersection.cpp 
    ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/intersection/RectangleIntersectionBuilder.cpp

    ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/linemerge/EdgeString.cpp 
	${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/linemerge/LineMergeDirectedEdge.cpp 
	${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/linemerge/LineMergeEdge.cpp 
	${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/linemerge/LineMergeGraph.cpp 
	${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/linemerge/LineMerger.cpp 
  ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/linemerge/LineSequencer.cpp
  
  ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/EdgeSetNoder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/ElevationMatrixCell.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/ElevationMatrix.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/LineBuilder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/MaximalEdgeRing.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/MinimalEdgeRing.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/OverlayNodeFactory.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/OverlayOp.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/PointBuilder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/PolygonBuilder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/snap/GeometrySnapper.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/snap/LineStringSnapper.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/snap/SnapOverlayOp.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/snap/SnapIfNeededOverlayOp.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/validate/FuzzyPointLocator.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/validate/OffsetPointGenerator.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/overlay/validate/OverlayResultValidator.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/polygonize/PolygonizeDirectedEdge.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/polygonize/PolygonizeEdge.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/polygonize/PolygonizeGraph.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/polygonize/Polygonizer.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/polygonize/EdgeRing.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/predicate/RectangleContains.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/predicate/RectangleIntersects.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/predicate/SegmentIntersectionTester.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/relate/EdgeEndBuilder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/relate/EdgeEndBundle.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/relate/EdgeEndBundleStar.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/relate/RelateComputer.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/relate/RelateNode.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/relate/RelateNodeFactory.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/relate/RelateNodeGraph.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/relate/RelateOp.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/sharedpaths/SharedPathsOp.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/union/CascadedPolygonUnion.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/union/CascadedUnion.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/union/PointGeometryUnion.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/union/UnaryUnionOp.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/valid/ConnectedInteriorTester.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/valid/ConsistentAreaTester.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/valid/IsValidOp.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/valid/QuadtreeNestedRingTester.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/valid/RepeatedPointTester.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/valid/SimpleNestedRingTester.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/valid/SweeplineNestedRingTester.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/valid/TopologyValidationError.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/valid/IndexedNestedRingTester.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/valid/IndexedNestedRingTester.h

          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/GeometryGraphOperation.cpp
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/operation/IsSimpleOp.cpp

          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/planargraph/DirectedEdge.cpp 
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/planargraph/DirectedEdgeStar.cpp 
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/planargraph/Edge.cpp 
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/planargraph/Node.cpp 
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/planargraph/NodeMap.cpp 
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/planargraph/PlanarGraph.cpp 
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/planargraph/Subgraph.cpp 
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/planargraph/algorithm/ConnectedSubgraphFinder.cpp

          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/precision/CommonBits.cpp 
	${CMAKE_CURRENT_SOURCE_DIR}/geos/src/precision/CommonBitsOp.cpp 
	${CMAKE_CURRENT_SOURCE_DIR}/geos/src/precision/CommonBitsRemover.cpp 
	${CMAKE_CURRENT_SOURCE_DIR}/geos/src/precision/EnhancedPrecisionOp.cpp 
	${CMAKE_CURRENT_SOURCE_DIR}/geos/src/precision/GeometryPrecisionReducer.cpp 
	${CMAKE_CURRENT_SOURCE_DIR}/geos/src/precision/MinimumClearance.cpp 
	${CMAKE_CURRENT_SOURCE_DIR}/geos/src/precision/PrecisionReducerCoordinateOperation.cpp 
  ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/precision/SimpleGeometryPrecisionReducer.cpp 
  
  ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/simplify/DouglasPeuckerLineSimplifier.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/simplify/DouglasPeuckerSimplifier.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/simplify/LineSegmentIndex.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/simplify/TaggedLineSegment.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/simplify/TaggedLineString.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/simplify/TaggedLineStringSimplifier.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/simplify/TaggedLinesSimplifier.cpp 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/simplify/TopologyPreservingSimplifier.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/triangulate/DelaunayTriangulationBuilder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/triangulate/IncrementalDelaunayTriangulator.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/triangulate/VoronoiDiagramBuilder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/triangulate/quadedge/LastFoundQuadEdgeLocator.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/triangulate/quadedge/LocateFailureException.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/triangulate/quadedge/QuadEdge.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/triangulate/quadedge/QuadEdgeLocator.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/triangulate/quadedge/QuadEdgeSubdivision.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/triangulate/quadedge/TrianglePredicate.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/triangulate/quadedge/TriangleVisitor.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/triangulate/quadedge/Vertex.cpp
  
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/util/Assert.cpp
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/util/GeometricShapeFactory.cpp
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/util/Interrupt.cpp
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/util/Profiler.cpp
          ${CMAKE_CURRENT_SOURCE_DIR}/geos/src/util/math.cpp
          
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/chain/MonotoneChain.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/chain/MonotoneChainBuilder.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/chain/MonotoneChainOverlapAction.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/chain/MonotoneChainSelectAction.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/strtree/AbstractNode.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/strtree/AbstractSTRtree.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/strtree/BoundablePair.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/strtree/Interval.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/strtree/ItemBoundable.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/strtree/SIRtree.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/strtree/STRtree.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/strtree/GeometryItemDistance.cpp

${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/quadtree/DoubleBits.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/quadtree/IntervalSize.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/quadtree/Key.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/quadtree/Node.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/quadtree/NodeBase.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/quadtree/Quadtree.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/quadtree/Root.cpp
 
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/bintree/Bintree.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/bintree/Interval.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/bintree/Bintree.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/bintree/Key.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/bintree/Node.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/bintree/NodeBase.cpp
${CMAKE_CURRENT_SOURCE_DIR}/geos/src/index/bintree/Root.cpp

)

FILE(GLOB src_algorithm geos/src/algorithm/**.cpp)
FILE(GLOB src_algorithm_distance geos/src/algorithm/distance/**.cpp)
FILE(GLOB src_algorithm_locate geos/src/algorithm/locate/**.cpp)
FILE(GLOB src_index_intervalrtree geos/src/index/intervalrtree/**.cpp)


set(HEADERS_GEOS
        ${CMAKE_CURRENT_SOURCE_DIR}/geos/include/geos.h
        ${CMAKE_CURRENT_SOURCE_DIR}/geos/capi/geos_c.h

)

ADD_LIBRARY(geos STATIC
  ${SRC_GEOS}
  ${HEADERS_GEOS}
  ${src_algorithm}
  ${src_algorithm_distance}
  ${src_algorithm_locate}
  ${src_index_intervalrtree}

          )

TARGET_INCLUDE_DIRECTORIES(geos PUBLIC $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/..>)

SOURCE_GROUP("src" FILES ${SRC_GEOS})
SOURCE_GROUP("include" FILES ${HEADERS_GEOS})

SET_TARGET_PROPERTIES(geos
                    PROPERTIES PUBLIC_HEADER "${headers}")
SET_TARGET_PROPERTIES(geos PROPERTIES
    ARCHIVE_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
    ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(geos PROPERTIES 
    LIBRARY_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
    LIBRARY_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(geos PROPERTIES
    RUNTIME_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
    RUNTIME_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
                    
IF(FLYWAVE_ENABLE_SOLUTION_FOLDERS)
  SET_TARGET_PROPERTIES(geos PROPERTIES FOLDER external)
ENDIF(FLYWAVE_ENABLE_SOLUTION_FOLDERS)
