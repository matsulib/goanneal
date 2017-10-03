##  50 USA Landmarks TSP Example

from [50 USA Landmarks](http://www.math.uwaterloo.ca/tsp/usa50/index.html)

According to the page, the optimal TSP tour (with the Google distances) is as follows:

```
 1. Grand Canyon National Park, Arizona                          (36.1070, -112.1130)
 2. Bryce Canyon National Park, Utah                             (37.5930, -112.1871)
 3. Craters of the Moon National Monument & Preserve, Idaho      (43.4166, -113.5166)
 4. Yellowstone National Park                                    (44.4621, -110.6424)
 5. Pikes Peak, Colorado 80809                                   (38.8409, -105.0423)
 6. Carlsbad Caverns National Park, New Mexico 88220             (32.1232, -104.5875)
 7. 300 Alamo Plaza, San Antonio, TX 78205                       (29.4260, -98.4861)
 8. Chickasha, OK                                                (34.4570, -97.0122)
 9. Toltec Mounds Archeological State Park, 490 Toltec Mounds Rd (34.6470, -92.0651)
10. Graceland, Elvis Presley Blvd, Memphis, TN 38116             (35.0477, -90.0260)
11. Vicksburg, MS                                                (32.3466, -90.8499)
12. French Quarter, New Orleans, LA                              (29.9584, -90.0644)
13. USS Alabama, 2703 Battleship Pkwy, Mobile, AL 36603          (30.6818, -88.0144)
14. Cape Canaveral, FL 32920                                     (28.3883, -80.6036)
15. Okefenokee Swamp Park, 5700 Okefenokee Swamp Park Rd         (31.0568, -82.2723)
16. Fort Sumter National Monument, South Carolina 29412          (32.7523, -79.8747)
17. 907 Lost World Rd, Lewisburg, WV 24901                       (37.8018, -80.4456)
18. 1000 N croatan hwy, Kill Devil Hills, NC 27948               (35.9082, -75.6757)
19. 3200 Mount Vernon Hwy, Mt Vernon, VA 22121                   (38.7293, -77.1074)
20. 1600 Pennsylvania Ave NW, Washington, DC 20500               (38.8977, -77.0365)
21. 100 State Cir, Annapolis, MD 21401                           (38.9788, -76.4910)
22. New Castle, DE 19720                                         (39.6582, -75.5623)
23. 200 Congress Pl, Cape May, NJ 08204                          (38.9318, -74.9242)
24. N 6th St & Market St, Philadelphia, PA 19106                 (39.9496, -75.1503)
25. Statue of Liberty National Monument, New York, NY 10004      (40.6892, -74.0445)
26. 351 Farmington Ave, Hartford, CT 06105                       (41.7668, -72.7012)
27. 44 Ochre Point Ave, Newport, RI 02840                        (41.4699, -71.2983)
28. Building 22, Charlestown Navy Yard, Charlestown, MA 02129    (42.3725, -71.0566)
29. Acadia National Park, Maine                                  (44.3386, -68.2733)
30. Bretton Woods, 99, 03575, Ski Area Rd, Jefferson, NH 03583   (44.2581, -71.4412)
31. 1611 Harbor Rd, Shelburne, VT 05482                          (44.4089, -73.2472)
32. 2211 Woodward Ave, Detroit, MI 48201                         (42.3876, -83.0849)
33. 4521 Spring Grove Ave, Cincinnati, OH 45232                  (39.1743, -84.5250)
34. Mammoth Cave National Park, Kentucky                         (37.1870, -86.1005)
35. 8538 W Baden Ave, West Baden Springs, IN 47469               (38.5667, -86.6175)
36. The Gateway Arch, St. Louis, MO 63102                        (38.6247, -90.1848)
37. 201 N 7th St #110, Lincoln, NE 68508                         (39.7975, -89.6462)
38. 5481 County Rd C, Spring Green, WI 53588                     (43.1410, -90.0705)
39. 200 Tower Ave, St Paul, MN 55111                             (44.8929, -93.1806)
40. 2300 Grand Ave, Des Moines, IA 50312                         (41.5832, -93.6485)
41. 320 S Esplanade St, Leavenworth, KS 66048                    (39.3172, -94.9095)
42. Ashfall Fossil Beds State Historical Park, 86930 517th Ave   (42.4250, -98.1586)
43. Mount Rushmore National Memorial, 13000 SD-244, Keystone     (43.8791, -103.4591)
44. 15550 ND-1804, Williston, ND 58801                           (48.0002, -104.0415)
45. Glacier National Park, Montana                               (48.7596, -113.7870)
46. Washington                                                   (46.5507, -119.4890)
47. Columbia River Gorge National Scenic Area, Hood River        (45.7116, -121.5196)
48. 1201 Mason St, San Francisco, CA 94108                       (37.7948, -122.4117)
49. San Andreas Fault, California                                (36.5761, -120.9876)
50. 118 Kingman Wash Access Rd, Temple Bar Marina, AZ 86443      (36.0161, -114.7377)
```

The above route is 22493 km HOWEVER I found some shorter routes. I'm not sure whether some new roads have been built or it's just a bug.

### Geographic Information

* distance_matrix.json (in meters)

The above landmarks' names are a little different from the original page. 
That's because Google Maps Distance Matrix API guessed the locations from the coordinates (longitude and latitude). 
