package database

import (
	"campsite/packages/event-service/internal/config"
	"campsite/packages/event-service/internal/service"
	"campsite/packages/event-service/internal/util"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

func getDbConnString(c *config.DbConfig) string {
	vals := getDbValues(c)
	var p []string
	for k, v := range vals {
		p = append(p, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(p, " ")
}

func getDbValues(c *config.DbConfig) map[string]string {
	p := map[string]string{}
	util.SetIfNotEmpty(p, "dbname", c.Name)
	util.SetIfNotEmpty(p, "host", c.Host)
	util.SetIfNotEmpty(p, "user", c.User)
	util.SetIfNotEmpty(p, "password", c.Password)
	util.SetIfNotEmpty(p, "port", c.Port)
	util.SetIfNotEmpty(p, "sslmode", c.SSLMode)
	return p
}

func NewDb(c *config.DbConfig) *gorm.DB {
	connStr := getDbConnString(c)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database")
	return db
}

// The same as NewDb but additionally migrates the database and creates
// mock data in the database.
func NewDevDb(c *config.DbConfig) *gorm.DB {
	db := NewDb(c)

	// Migrate the database.
	if err := db.AutoMigrate(&service.Event{}, &service.Speaker{}, &service.Session{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	log.Println("Auto migrated database")

	// Create a mock event in the database.
	mockEvent := getMockEvent()
	if err := db.Create(&mockEvent).Error; err != nil {
		// The error will likely occur because we already created it already,
		// that is, the primary keys we set up above already exists.
		// We can ignore this.
		log.Printf("Failed to create mock event in database: %v", err)
	}
	log.Println("Created mock event in database")

	// Create a mock OpenCloudConf event in the database.
	mockOpenCloudConfEvent := getMockOpenCloudConfEvent()
	if err := db.Create(&mockOpenCloudConfEvent).Error; err != nil {
		// The error will likely occur because we already created it already,
		// that is, the primary keys we set up above already exists.
		// We can ignore this.
		log.Printf("Failed to create mock OpenCloudConf event in database: %v", err)
	}
	log.Println("Created OpenCloudConf mock event in database")

	return db
}

func getMockEvent() service.Event {
	now := time.Now()
	later := now.Add(time.Hour * 1)
	evenLater := later.Add(time.Hour * 2)
	evenEvenLater := evenLater.Add(time.Hour * 4)
	evenEvenEvenLater := evenEvenLater.Add(time.Hour * 22)
	evenEvenEvenEvenLater := evenEvenLater.Add(time.Hour * 2)

	spk1 := service.Speaker{
		ID:       "9c08fbf8-160b-4a86-9981-aeddf4e3798e",
		Name:     "Spencer Waldron",
		Bio:      "Global Communications Director for Prezi Video - a tool for online classes and learning. Opinions are my own.",
		Headline: "Head of Remote",
		Photo:    "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a8599ecda6125a34ad3dc_Spencer%20Waldron.jpeg",
	}

	spk2 := service.Speaker{
		ID:       "361655d7-3034-426b-924f-589c79533650",
		Name:     "Iwo Szapar",
		Bio:      "Loop Team is a virtual office that brings the best parts of an office environment to distributed teams - stay in the loop.",
		Headline: "CEO of Remote-how",
		Photo:    "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3ed56e8ee3ae185f0500e5_Iwo%20Szapar.jpeg",
	}

	spk3 := service.Speaker{
		ID:       "1c2d5f82-a5ee-4cf2-80ea-b134c0f6d969",
		Name:     "Mike Adams",
		Bio:      "Loop Team is a virtual office that brings the best parts of an office environment to distributed teams - stay in the loop.",
		Headline: "CEO of Grain",
		Photo:    "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3f7be5fbaf556a9447ed80_Mike%20Adams.jpeg",
	}

	sess1 := service.Session{
		ID:          "be13940b-c7ba-4f97-bdab-b4a47b11ffed",
		Name:        "How to build and maintain great company culture remotely",
		StartDate:   &now,
		EndDate:     &later,
		Description: "On the other hand, we denounce with righteous indignation and dislike men who are so beguiled and demoralized by the charms of pleasure of the moment, so blinded by desire, that they cannot foresee the pain and trouble that are bound to ensue; and equal blame belongs to those who fail in their duty through weakness of will, which is the same as saying through shrinking from toil and pain. These cases are perfectly simple and easy to distinguish.",
		Url:         "https://google.com",
		EventID:     "ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
		Speakers:    []service.Speaker{spk1, spk2},
	}

	sess2 := service.Session{
		ID:          "8b83ebd8-b7f1-4ef3-a141-2049ed59232f",
		Name:        "Connecting your workforce through your company values",
		StartDate:   &evenLater,
		EndDate:     &evenEvenLater,
		Description: "In a free hour, when our power of choice is untrammelled and when nothing prevents our being able to do what we like best, every pleasure is to be welcomed and every pain avoided. But in certain circumstances and owing to the claims of duty or the obligations of business it will frequently occur that pleasures have to be repudiated and annoyances accepted. The wise man therefore always holds in these matters to this principle of selection: he rejects pleasures to secure other greater pleasures, or else he endures pains to avoid worse pains.",
		Url:         "https://google.com",
		EventID:     "ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
		Speakers:    []service.Speaker{spk2, spk1},
	}

	sess3 := service.Session{
		ID:          "238f8433-483c-4266-b687-6b8a81ccc39e",
		Name:        "Connecting your workforce through your company values",
		StartDate:   &evenEvenEvenLater,
		EndDate:     &evenEvenEvenEvenLater,
		Description: "In a free hour, when our power of choice is untrammelled and when nothing prevents our being able to do what we like best, every pleasure is to be welcomed and every pain avoided. But in certain circumstances and owing to the claims of duty or the obligations of business it will frequently occur that pleasures have to be repudiated and annoyances accepted. The wise man therefore always holds in these matters to this principle of selection: he rejects pleasures to secure other greater pleasures, or else he endures pains to avoid worse pains.",
		Url:         "https://google.com",
		EventID:     "ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
		Speakers:    []service.Speaker{spk3, spk2},
	}

	address := "San Francisco, CA"
	event := service.Event{
		ID:              "ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
		Name:            "BigDataCamp LA 2020",
		Description:     "BigDataCamp is an unconference where early adopters of BigData technologies, such as Hadoop, exchange ideas. With the rapid change occurring in the industry, we need a place where we can meet to share our experiences, challenges and solutions. At BigDataCamp, you are encouraged to share your thoughts in several open discussions, as we strive for the advancement of BigData. Data Scientists, Developers, IT professionals, users and vendors are all encouraged to participate.",
		RegistrationUrl: "https://www.eventbrite.com/e/redis-day-london-2019-registration-71402886957#",
		StartDate:       &now,
		EndDate:         &later,
		Photo:           "https://events.redislabs.com/wp-content/uploads/2020/04/redisconf2020-hero-m-4.png",
		OrganizerName:   "Tim Apple",
		Address:         &address,
		Sessions:        []service.Session{sess1, sess2, sess3},
	}
	return event
}

func getMockOpenCloudConfEvent() *service.Event {
	spkRandy := newSpeaker("Randy Bias", "", "CloudScaling", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3cfb15cabbb341f90be301_Cheryl%20Crane.jpeg")
	spkGreg := newSpeaker("Greg Smith", "", "Eucalyptus", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg")
	spkJoe := newSpeaker("Joe Arnold", "", "Apple", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3cf73b49bf9b6356f235a8_Katherine_Zaleski_Remote%20Future%20Summit.jpg")
	spkMark := newSpeaker("Mark Hinkle", "", "Cloudstack.org", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a825f7423af9606c5e31d_Shiran%20Yaroslavsky_Remote%20Future%20Summit.jpeg")
	spkDave := newSpeaker("Dave Nielsen", "", "Traceable", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f6215f19526a65368fc86a0_Adam%20Hickman.jpeg")
	spkDiane := newSpeaker("Diane Mueller", "", "Tesla", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a89d7e5ef2b48c706639d_Chase%20Warrington.jpg")
	spkGordon := newSpeaker("Gordon Haff", "", "Amazon", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f47fc750ee8863f7b7a598a_Brendan%20O%27Leary.jpeg")
	spkAdrian := newSpeaker("Adrian Cole", "", "Google", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f624fecb46e2612148cfee3_Liam%20Martin.jpeg")
	spkJaisen := newSpeaker("Jaisen Mathai", "", "Uber", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f5d10478b598e75dd657788_Katrina%20Collier_Remote%20Future%20Summit%202020.jpg")

	// Pacific Time (PDT)
	loc := time.FixedZone("UTC-7", -7*60*60)
	eventStartDate := time.Date(2012, time.May, 1, 9, 0, 0, 0, loc)
	eventEndDate := time.Date(2012, time.May, 2, 5, 0, 0, 0, loc)

	// Day 1
	ss1 := newSession("Best of Breed: Why Open Clouds are Better", eventStartDate, time.Minute*30, "", "", *spkDave)
	ss2 := newSession("The State of the Open Cloud", *ss1.EndDate, time.Minute*45, "", "", *spkRandy)
	ss3 := newSession("Open Cloud vs. Open Source: What's the difference?", *ss2.EndDate, time.Minute*45, "", "", *spkGordon, *spkDiane)
	ss4 := newSession("Open Cloud APIs - Why All the Fuss? Can an API be THAT important?", *ss3.EndDate, time.Minute*45, "", "", *spkAdrian)
	ss5 := newSession("OpenStack Workshop Part 1", *ss4.EndDate, time.Minute*90, "", "", *spkJoe)
	ss6 := newSession("Real Key to Open Cloud: Building in Cloud Application Portability", *ss5.EndDate, time.Minute*45, "", "", *spkGordon)
	ss7 := newSession("OpenPaaS & Open Eucalyptus", *ss6.EndDate, time.Minute*30, "", "", *spkGreg)
	ss8 := newSession("OpenStack Workshop Part 2", *ss7.EndDate, time.Minute*105, "", "", *spkJoe)
	ss9 := newSession("Avoiding Cloud-Lock-In", *ss8.EndDate, time.Minute*60, "", "", *spkMark)
	ss10 := newSession("Application Portability in the Cloud", *ss9.EndDate, time.Minute*45, "", "", *spkDiane)
	// Day 2
	ss11 := newSession("The Cloud: Open for Business", eventStartDate.Add(time.Hour*24), time.Minute*30, "", "", *spkDave)
	ss12 := newSession("Understand cloud begins with the public cloud", *ss11.EndDate, time.Minute*45, "", "", *spkAdrian)
	ss13 := newSession("OpenPhoto, Personal Clouds and Why it Matters", *ss12.EndDate, time.Minute*30, "", "", *spkJaisen)
	ss14 := newSession("Scripts, Images and PaaS, Oh My!", *ss13.EndDate, time.Minute*30, "", "", *spkRandy)
	// The next 2 sessions are being held simultaneously.
	ss15 := newSession("DevOps/PaaS Workshop: Open Source Multi-cloud Application Management", *ss14.EndDate, time.Minute*300, "", "", *spkJoe)
	ss16 := newSession("Developer/PaaS Workshop: Stackato/Cloud Foundry", *ss14.EndDate, time.Minute*300, "", "", *spkDave)
	ss17 := newSession("OpenShift + OpenStack + Fedora = Awesome!", *ss15.EndDate, time.Minute*45, "", "", *spkMark, *spkGreg)

	eventAddress := "Mountain View, CA"
	event := newEvent(
		"OpenCloudConf",
		"Something new this year that we picked up along they way, as we know not every can and also not everyone wants to travel so that is why we are making Shift Dev a Hybrid event. What does that mean? Well for starters all talks will be professionally streamed so anyone from around the world can tune in. Second you will be able to chat remotely with everyone at the event - both live and remote, and lastly you will be able to remotely ask all the speakers direct questions via our Remote only AMA with each of our speakers! So in Short you get to meet new people, listen to all the talks and talk directly to the speakers themselves!",
		"https://apple.com",
		eventStartDate,
		eventEndDate,
		"https://azurecomcdn.azureedge.net/cvt-535bb5cd07ecbbf429a895be32834e506bd399cd4f6a8e4bc33a27bd5ffee836/images/page/services/devops/hero-images/index-hero.jpg",
		"Dave Nielsen",
		&eventAddress,
		// Day 1
		ss1, ss2, ss3, ss4, ss5, ss6, ss7, ss8, ss9, ss10,
		// Day 2
		ss11, ss12, ss13, ss14, ss15, ss16, ss17)

	return event
}

func newEvent(name string, description string, registrationUrl string, startDate time.Time, endDate time.Time, photo string, organizerName string, address *string, sessions ...service.Session) *service.Event {
	return &service.Event{
		ID:              uuid.New().String(),
		Name:            name,
		Description:     description,
		RegistrationUrl: registrationUrl,
		StartDate:       &startDate,
		EndDate:         &endDate,
		Photo:           photo,
		OrganizerName:   organizerName,
		Address:         address,
		Sessions:        sessions,
	}
}

func newSession(name string, startDate time.Time, duration time.Duration, description string, url string, speakers ...service.Speaker) service.Session {
	endDate := startDate.Add(duration)
	return service.Session{
		ID:          uuid.New().String(),
		Name:        name,
		StartDate:   &startDate,
		EndDate:     &endDate,
		Description: description,
		Url:         url,
		Speakers:    speakers,
	}
}

func newSpeaker(name string, bio string, headline string, photo string) *service.Speaker {
	return &service.Speaker{
		ID:       uuid.New().String(),
		Name:     name,
		Bio:      bio,
		Headline: headline,
		Photo:    photo,
	}
}
