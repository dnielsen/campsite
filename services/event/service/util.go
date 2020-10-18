package service

import "github.com/dnielsen/campsite/pkg/model"

func GetUniqueSpeakersFromSessions(sessions []model.Session) []model.Speaker {
	// Iterate through each session and each speaker of the sessions
	// and create an array of unique speakers.
	// The key of the map is a speaker id and the value is a speaker.
	spkMap := make(map[string]model.Speaker)
	for _, sess := range sessions {
		for _, spk := range sess.Speakers {
			spkMap[spk.ID] = spk
		}
	}

	// Get the values (Speakers) of the map.
	var uniqSpks []model.Speaker
	for _, spk := range spkMap {
		uniqSpks = append(uniqSpks, spk)
	}

	return uniqSpks
}
