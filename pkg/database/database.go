package database

import (
	"fmt"
	"github.com/dnielsen/campsite/pkg/config"
	"github.com/dnielsen/campsite/pkg/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

func setIfNotEmpty(m map[string]string, key, val string) {
	if val != "" {
		m[key] = val
	}
}

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
	setIfNotEmpty(p, "dbname", c.Name)
	setIfNotEmpty(p, "host", c.Host)
	setIfNotEmpty(p, "user", c.User)
	setIfNotEmpty(p, "password", c.Password)
	setIfNotEmpty(p, "port", c.Port)
	setIfNotEmpty(p, "sslmode", c.SSLMode)
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


// The same as NewDb but additionally migrates the database and might create
// mock data in the database.
func NewDevDb(c *config.Config) *gorm.DB {
	db := NewDb(&c.Db)
	// Migrate the database.
	if err := db.AutoMigrate(&model.Event{}, &model.Speaker{}, &model.Session{}, &model.User{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	} else {
		log.Println("Auto migrated database")
	}
	// Create events, sessions, speakers, users in db.
	if c.Dev.MockEnabled {
		createMockDataInDb(db, c)
		log.Printf("Created mock data in database")
	}

	return db
}

func createMockDataInDb(db *gorm.DB, c *config.Config) {
	// Create the only user in the database that has permissions to create/edit/delete stuff.
	u, err := newUser(c.Admin.Email, c.Admin.Password)
	if err != nil {
		log.Printf("Failed to auto create user: %v", err)
	} else {
		if err := db.Create(&u).Error; err != nil {
			log.Printf("Failed to create mock user in database: %v", err)
		} else {
			log.Printf("Created admin user %v:%v in database", c.Admin.Email, c.Admin.Password)
		}
	}
	// Create a mock event in the database.
	mockEvent := newMockEvent()
	if err := db.Create(&mockEvent).Error; err != nil {
		// The error will likely occur because we already created it already,
		// that is, the primary keys we set up above already exists.
		// We can ignore this.
		log.Printf("Failed to create mock event in database: %v", err)
	} else {
		log.Println("Created mock event in database")
	}
	// Create a mock OpenCloudConf event in the database.
	mockOpenCloudConfEvent := newMockOpenCloudConfEvent()
	if err := db.Create(&mockOpenCloudConfEvent).Error; err != nil {
		// The error will likely occur because we already created it already,
		// that is, the primary keys we set up above already exists.
		// We can ignore this.
		log.Printf("Failed to create mock OpenCloudConf event in database: %v", err)
	} else {
		log.Println("Created OpenCloudConf mock event in database")
	}
}

func newUser(email, password string) (*model.User, error) {
	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:           uuid.New().String(),
		Email:        email,
		PasswordHash: string(passwordHashBytes),
	}, nil
}

func newMockEvent() model.Event {
	now := time.Now()
	later := now.Add(time.Hour * 5)
	evenLater := later.Add(time.Hour * 5)

	spk1 := newSpeaker("Spencer Waldron",
		"Global Communications Director for Prezi Video - a tool for online classes and learning. Opinions are my own.",
		"Head of Remote",
		"https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a8599ecda6125a34ad3dc_Spencer%20Waldron.jpeg")

	spk2 := newSpeaker("Iwo Szapar",
		"Loop Team is a virtual office that brings the best parts of an office environment to distributed teams - stay in the loop.",
		"CEO of Remote-how",
		"https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3ed56e8ee3ae185f0500e5_Iwo%20Szapar.jpeg")

	spk3 := newSpeaker("Mike Adams",
		"Global Communications Director for Prezi Video - a tool for online classes and learning. Opinions are my own.",
		"CEO of Grain",
		"https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3f7be5fbaf556a9447ed80_Mike%20Adams.jpeg")

	ss1 := newSession("How to build and maintain great company culture remotely",
		now,
		time.Minute*180,
		"On the other hand, we denounce with righteous indignation and dislike men who are so beguiled and demoralized by the charms of pleasure of the moment, so blinded by desire, that they cannot foresee the pain and trouble that are bound to ensue; and equal blame belongs to those who fail in their duty through weakness of will, which is the same as saying through shrinking from toil and pain. These cases are perfectly simple and easy to distinguish.",
		"https://google.com",
		spk1,
		spk2)

	ss2 := newSession("Connecting your workforce through your company values",
		later,
		time.Minute*90,
		"In a free hour, when our power of choice is untrammelled and when nothing prevents our being able to do what we like best, every pleasure is to be welcomed and every pain avoided. But in certain circumstances and owing to the claims of duty or the obligations of business it will frequently occur that pleasures have to be repudiated and annoyances accepted. The wise man therefore always holds in these matters to this principle of selection: he rejects pleasures to secure other greater pleasures, or else he endures pains to avoid worse pains.",
		"https://google.com",
		spk1)

	ss3 := newSession("Connecting your workforce through your company values",
		evenLater,
		time.Minute*45,
		"In a free hour, when our power of choice is untrammelled and when nothing prevents our being able to do what we like best, every pleasure is to be welcomed and every pain avoided. But in certain circumstances and owing to the claims of duty or the obligations of business it will frequently occur that pleasures have to be repudiated and annoyances accepted. The wise man therefore always holds in these matters to this principle of selection: he rejects pleasures to secure other greater pleasures, or else he endures pains to avoid worse pains.",
		"https://google.com",
		spk3, spk2)

	address := "San Francisco, CA"
	event := newEvent("BigDataCamp",
		"BigDataCamp is an unconference where early adopters of BigData technologies, such as Hadoop, exchange ideas. With the rapid change occurring in the industry, we need a place where we can meet to share our experiences, challenges and solutions. At BigDataCamp, you are encouraged to share your thoughts in several open discussions, as we strive for the advancement of BigData. Data Scientists, Developers, IT professionals, users and vendors are all encouraged to participate.",
		"https://www.eventbrite.com/e/redis-day-london-2019-registration-71402886957#",
		now,
		now.Add(time.Hour*10+time.Minute*45),
		"https://events.redislabs.com/wp-content/uploads/2020/04/redisconf2020-hero-m-4.png",
		"Tim Apple",
		&address,
		ss1, ss2, ss3)

	return event
}

func newMockOpenCloudConfEvent() model.Event {
	spkRandy := newSpeaker("Randy Bias", "He has worked for over 20 years as a developer, team leader and founder & CEO. After seeing his company through to an acquisition by Oracle, he now leads Product Development for Oracle's API portfolio. Founder of cesko.digital.", "CloudScaling", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3cfb15cabbb341f90be301_Cheryl%20Crane.jpeg")
	spkGreg := newSpeaker("Greg Smith", "David is a Google Developer Expert for Android. He leads his startup and also works as a Senior Android Developer at JLL. He loves open-source, Tesla, and LARP.", "Eucalyptus", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg")
	spkJoe := newSpeaker("Joe Arnold", "Magda Miu is an enthusiastic and innovative Squad Lead Developer at Orange and Google Developer Expert for Android with more than 10 years experience in software development.", "Apple", "https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3cf73b49bf9b6356f235a8_Katherine_Zaleski_Remote%20Future%20Summit.jpg")
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
	ss1 := newSession("Best of Breed: Why Open Clouds are Better", eventStartDate, time.Minute*30, "Navigation is one the most fundamental element to any Android application. We will talk about how to use Android Navigation component to get this crucial aspect of our apps right, maintainable and easy to reason with.\n\nWe will talk about how the navigation component allows us to isolate the navigation from rest of the application logic and gives a nice overview of the whole application navigation in one single graph.", "https://apple.com", spkDave)
	ss2 := newSession("The State of the Open Cloud", *ss1.EndDate, time.Minute*45, "In 90 minutes we'll debunk the myth of iOS being secure-by-default, walk through the various techniques of penetration testing, try out a plethora of tools for security testing and learn how to make our systems as robust as possible.", "https://apple.com", spkRandy)
	ss3 := newSession("Open Cloud vs. Open Source: What's the difference?", *ss2.EndDate, time.Minute*45, "In this talk, we will take a deep dive into the Android keystore system, certificates and signatures, and go over key points necessary for any application's long and productive life. Also, we will explore some security tips and tricks that will help ensure your app is safe to use, even if the users are faced with its evil twin.", "https://apple.com", spkGordon, spkDiane)
	ss4 := newSession("Open Cloud APIs - Why All the Fuss? Can an API be THAT important?", *ss3.EndDate, time.Minute*45, "", "https://apple.com", spkAdrian)
	ss5 := newSession("OpenStack Workshop Part 1", *ss4.EndDate, time.Minute*90, "", "", spkJoe)
	ss6 := newSession("Real Key to Open Cloud: Building in Cloud Application Portability", *ss5.EndDate, time.Minute*45, "", "", spkGordon)
	ss7 := newSession("OpenPaaS & Open Eucalyptus", *ss6.EndDate, time.Minute*30, "", "", spkGreg)
	ss8 := newSession("OpenStack Workshop Part 2", *ss7.EndDate, time.Minute*105, "", "", spkJoe)
	ss9 := newSession("Avoiding Cloud-Lock-In", *ss8.EndDate, time.Minute*60, "", "", spkMark)
	ss10 := newSession("Application Portability in the Cloud", *ss9.EndDate, time.Minute*45, "", "", spkDiane)
	// Day 2
	ss11 := newSession("The Cloud: Open for Business", eventStartDate.Add(time.Hour*24), time.Minute*30, "", "", spkDave)
	ss12 := newSession("Understand cloud begins with the public cloud", *ss11.EndDate, time.Minute*45, "", "", spkAdrian)
	ss13 := newSession("OpenPhoto, Personal Clouds and Why it Matters", *ss12.EndDate, time.Minute*30, "", "", spkJaisen)
	ss14 := newSession("Scripts, Images and PaaS, Oh My!", *ss13.EndDate, time.Minute*30, "", "", spkRandy)
	// The next 2 sessions are being held simultaneously.
	ss15 := newSession("DevOps/PaaS Workshop: Open Source Multi-cloud Application Management", *ss14.EndDate, time.Minute*300, "", "", spkJoe)
	ss16 := newSession("Developer/PaaS Workshop: Stackato/Cloud Foundry", *ss14.EndDate, time.Minute*300, "", "", spkDave)
	ss17 := newSession("OpenShift + OpenStack + Fedora = Awesome!", *ss15.EndDate, time.Minute*45, "", "", spkMark, spkGreg)

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

func newEvent(name string, description string, registrationUrl string, startDate time.Time, endDate time.Time, photo string, organizerName string, address *string, sessions ...model.Session) model.Event {
	return model.Event{
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

func newSession(name string, startDate time.Time, duration time.Duration, description string, videoUrl string, speakers ...model.Speaker) model.Session {
	endDate := startDate.Add(duration)
	return model.Session{
		ID:          uuid.New().String(),
		Name:        name,
		StartDate:   &startDate,
		EndDate:     &endDate,
		Description: description,
		VideoUrl:    videoUrl,
		Speakers:    speakers,
	}
}

func newSpeaker(name string, bio string, headline string, photo string) model.Speaker {
	return model.Speaker{
		ID:       uuid.New().String(),
		Name:     name,
		Bio:      bio,
		Headline: headline,
		Photo:    photo,
	}
}