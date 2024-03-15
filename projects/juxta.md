---
id: 1710461607-IEOE
aliases:
  - juxta-overview
tags: []
area: ""
project: ""
---

# juxta-overview

### source code: private github organization

Juxta is a ride-hailing app solution for hospitals that aims to improve delay time 
and efficiency of room-to-room patient transports in hospitals. It enables nurses
or doctors to put in requests for patients to be transported from an origin such as
the ER to a destination such as a room.

It aims to facilitate this in three sections: A navigation engine, an indoor 
positioning system, a machine learning floorplan digitization suite, and a mobile app.

Firebase is used for many backend needs. Cloud storage is used to store custom floormaps
with paths and rooms information. Firestore is used for every database need. Cloud functions
are used to expose the navigation engine as an api endpoint.

## Juxta Mobile App

Juxta's mobile app is completely written in typescript using React Native for
cross-platform compatibility. The implementation is almost fully complete excepting
some pre-shipping debugging and implementation of a brief list of features.

Notably, Expo managed workflow is used within the app for build efficiency with
noticably less native compatibility configuration issues. EAS cloud builds combined
with Expo Go builds make it very easy to develop on every platform. 

UI is streamlined using eva design system with the help of UI Kitten. RN reanimated
is used throughout for UI animations. RN maps with custom Google Maps styling is used,
which is the focal point of the application. GeoJson formatted custom map data is
overlaid on top of this map to visualize indoor maps and display navigation routing
that queries a custom api endpoint of the Juxta navigation engine.

![[juxta-nav.png]]

Indoor positioning system is partly integrated into the mobile app in the form of 
a synthesis use of gps geolocation services and another custom api endpoint for the
Juxta indoor positioning system, which gets streamed sensor data. Software smoothing 
is used to snap user location to a point on a path on the map.

In summary, Juxta mobile app combines the positioning tools in a user interface.

### TODO:
- [ ] Push notifications
- [ ] EAS build testing

## Juxta Navigation Engine

Navigation engine is written in JavaScript as a set of Firebase cloud functions and 
provides custom pathfinding on any floorplan. The navigation engine
is almost complete as well with the exception of building-to-building routing. It
can, however, do room-to-room and floor-to-floor routing.

Several data structure and algorithms are at play, while both initializing a map
of nodes for navigating and providing pathfinding. While initializing a map, 
a node graph is used along with the Bentley-Ottman algorithm to find intersections
of seperate paths on raw GeoJson formatted data that is uploaded to cloud storage.
This initial pass generates and serializes the node graph into a JSON.

Pathfinding is simply done by fetching the serialized JSON of the node graph and
using NBA* bi-directional pathfinding algorithm. NBA* outperforms Dijkstra's and 
A*. A quadtree is used to snap user's location on the node graph as it is nearly impossible
that the streamed user location of longitude and latitude matches up exactly with a 
node's location. Hence, the resolution of the quadtree is high for accuracy.
If a valid path exists for a given origin (user location) and a destination the 
navigation engine responds in a JSON containing information for eta, legs, and nodes to 
visit. 

The node-graph generated can be inspected using a set of custom graph tools that can
visualize a node map of a floorplan of a hospital in this case. It is written in HTML
and JavaScript and used with Live Server on localhost.

The quadtree and node-graph setup is paralleled in the mobile app for granular or
smooth routing or navigation experience.

### TODO:
- [ ] Migration to Go 

## Juxta Machine Learning

Machine learning suite is written in Python and is currently just an initial attempt at 
what would be a huge solution in indoor positioning.

MRCNN is used for instance segmentation of rooms given an image of a floorplan.
The model that has been trained is quite limited in its capability to detect rooms
and put boundaries to them as the dataset that was used to train it is only a segmented
map of Northwestern University's Tech building's first floor. 

### TODO:
- [ ] Improve CNN design
- [ ] Extend dataset w/ COCO
- [ ] Integrate admin UI for edits  

## Juxta Indoor Positioning

Juxta indoor positioning is currently a solution written in Go. It is a synthesis 
of existing techniques and our software smoothing for better accuracy. It is yet
incomplete as well and needs to be finished before shipping.

It implements sensor fusion, a pedometer to determine step size and distance travelled,
and fingerprinting to boost accuracy of location. 

### TODO:
- [ ] Test out ML integration 
- [ ] Deploy in a Docker container 

