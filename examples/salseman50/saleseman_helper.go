// +build ignore

package main

import (
	"math"
	"math/rand"
)

// shuffle a Slice
func shuffle(data []string) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

// Convert degrees to radians
func radians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

// Calculates distance between two latitude-longitude coordinates.
func distance(a [2]float64, b [2]float64) float64 {
	R := 3963.0 // radius of Earth (miles)
	lat1, lon1 := radians(a[0]), radians(a[1])
	lat2, lon2 := radians(b[0]), radians(b[1])
	return math.Acos(math.Sin(lat1)*math.Sin(lat2)+
		math.Cos(lat1)*math.Cos(lat2)*math.Cos(lon1-lon2)) * R
}

type Tuple [2]float64
type DMap map[string]float64

type America struct {
	Cities         map[string]Tuple
	DistanceMatrix map[string]DMap
}

func NewAmerica() *America {
	a := new(America)
	a.Cities = map[string]Tuple{
		"Grand Canyon, AZ":                      {36.106965, 112.112997},
		"Bryce Canyon, UT":                      {37.593038, 112.187090},
		"Craters of the Moon, ID":               {43.416650, 113.516650},
		"Yellowstone, WY":                       {44.462085, 110.642441},
		"Pikes Peak, CO":                        {38.840871, 105.042260},
		"Carlsbad Caverns, NM":                  {32.123169, 104.587450},
		"The Alamo, TX":                         {29.425967, 98.486142},
		"Chickasaw, OK":                         {34.457043, 97.012213},
		"Toltec Mounds, AR":                     {34.647037, 92.065143},
		"Graceland, TN":                         {35.047691, 90.026049},
		"Vicksburg, MS":                         {32.346550, 90.849850},
		"French Quarter, New Orleans, LA":       {29.958443, 90.064411},
		"USS Alabama, AL":                       {30.681803, 88.014426},
		"Cape Canaveral, FL":                    {28.388333, 80.603611},
		"Okefenokee Swamp, GA":                  {31.056794, 82.272327},
		"Fort Sumter, SC":                       {32.752348, 79.874692},
		"Lost World Caverns, WV":                {37.801788, 80.445630},
		"Wright Brothers Visitor Center, NC":    {35.908226, 75.675730},
		"Mount Vernon, VA":                      {38.729314, 77.107386},
		"White House, DC":                       {38.897676, 77.036530},
		"Maryland State House, MD":              {38.978828, 76.490974},
		"New Castle Historic District, DE":      {39.658242, 75.562335},
		"Congress Hall, Cape May, NJ":           {38.931843, 74.924184},
		"Liberty Bell, PA":                      {39.949610, 75.150282},
		"Statue of Liberty, NY":                 {40.689249, 74.044500},
		"Mark Twain House, Hartford, CT":        {41.766759, 72.701173},
		"The Breakers, Newport, RI":             {41.469858, 71.298265},
		"USS Constitution, Boston, MA":          {42.372470, 71.056575},
		"Acadia National Park, ME":              {44.338556, 68.273335},
		"Mount Washington, Bretton Woods, NH":   {44.258120, 71.441189},
		"Shelburne Farms, VT":                   {44.408948, 73.247227},
		"Olympia Entertainment, Detroit, MI":    {42.387579, 83.084943},
		"Spring Grove Cemetery, Cincinnati, OH": {39.174331, 84.524997},
		"Mammoth Cave National Park, KY":        {37.186998, 86.100528},
		"West Baden Springs Hotel, IN":          {38.566697, 86.617524},
		"Gateway Arch, St. Louis, MO":           {38.624691, 90.184776},
		"Lincoln Visitor Center, IL":            {39.797519, 89.646184},
		"Taliesin, WI":                          {43.141031, 90.070467},
		"Fort Snelling, MN":                     {44.892850, 93.180627},
		"Terrace Hill, IA":                      {41.583218, 93.648542},
		"C. W. Parker Carousel Museum, KS":      {39.317245, 94.909536},
		"Ashfall Fossil Bed, NE":                {42.425000, 98.158611},
		"Mount Rushmore, SD":                    {43.879102, 103.459067},
		"Fort Union Trading Post, ND":           {48.000160, 104.041483},
		"Glacier National Park, MT":             {48.759613, 113.787023},
		"Hanford Site, WA":                      {46.550684, 119.488974},
		"Columbia River Gorge, OR":              {45.711564, 121.519633},
		"Cable Car Museum, San Francisco, CA":   {37.794781, 122.411715},
		"San Andreas Fault, CA":                 {36.576088, 120.987632},
		"Hoover Dam, NV":                        {36.016066, 114.737732}}

	// create a distance matrix
	a.DistanceMatrix = map[string]DMap{}
	for ka, va := range a.Cities {
		a.DistanceMatrix[ka] = DMap{}
		for kb, vb := range a.Cities {
			if kb == ka {
				a.DistanceMatrix[ka][kb] = 0.0
			} else {
				a.DistanceMatrix[ka][kb] = distance(va, vb)
			}
		}
	}
	return a
}

func (self *America) CitiesKeys() []string {
	ks := []string{}
	for k, _ := range self.Cities {
		ks = append(ks, k)
	}
	return ks
}
