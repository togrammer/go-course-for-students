package tagcloud

import "math"

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	StatSlice []TagStat
}

// TagStat represents statistics regarding single tag
type TagStat struct {
	Tag             string
	OccurrenceCount int
}

// New should create a valid TagCloud instance
func New() TagCloud {
	OurSlice := []TagStat{}
	return TagCloud{StatSlice: OurSlice}
}

// AddTag should add a tag to the cloud if it wasn't present and increase tag occurrence count
// thread-safety is not needed
func (cloud *TagCloud) AddTag(tag string) {
	for i, v := range cloud.StatSlice {
		if v.Tag == tag {
			cloud.StatSlice[i].OccurrenceCount++
			for j := i - 1; j >= 0; j-- {
				if cloud.StatSlice[j].OccurrenceCount < cloud.StatSlice[j+1].OccurrenceCount {
					cloud.StatSlice[j], cloud.StatSlice[j+1] = cloud.StatSlice[j+1], cloud.StatSlice[j]
				}
			}
			return
		}
	}
	cloud.StatSlice = append(cloud.StatSlice, TagStat{Tag: tag, OccurrenceCount: 1})
}

// TopN should return top N most frequent tags ordered in descending order by occurrence count
// if there are multiple tags with the same occurrence count then the order is defined by implementation
// if n is greater that TagCloud size then all elements should be returned
// thread-safety is not needed
// there are no restrictions on time complexity
func (cloud *TagCloud) TopN(n int) []TagStat {
	// TODO: Implement this
	return cloud.StatSlice[0:int(math.Min(float64(n), float64(len(cloud.StatSlice))))]
}
