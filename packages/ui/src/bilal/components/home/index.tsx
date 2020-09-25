import React, { Fragment } from "react";
import { EventDetails } from "../../common/interfaces";
import SpeakerList from "../speaker/SpeakerList";
import SessionSchedule from "../session/SessionSchedule";
import util from "../../common/util";
import { Tab, Tabs, TabList, TabPanel } from "react-tabs";

// import styled component
import * as s from "../../styles/homePageStyles";
import * as g from "../../styles/globalStyles";

function HomePage() {
  return (
    <Fragment>
      <g.Container>
        <s.BannerSection>
          <p>23 - 24 September, online.</p>
          <h1>
            Remote Future
            <br /> Summit 2020
          </h1>
          <h3>Remote work can be difficult.</h3>
          <h3>We're here to talk about it.</h3>
        </s.BannerSection>
        <s.SelectSeatButton>
          <a href="/">Save your seat!</a>
        </s.SelectSeatButton>
        <s.LearnFromBestLogos>
          <p>Learn from the best:</p>
          <s.BrandImages>
            <img
              src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f4e5b56269716862f12fd40_slido_black.png"
              alt="brand-logo"
            />
            <img
              src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f354394d9a1463ff9f2f0ff_SmartRecruiters.png"
              alt="brand-logo"
            />
            <img
              src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f4e5b56269716862f12fd40_slido_black.png"
              alt="brand-logo"
            />
            <img
              src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f354394d9a1463ff9f2f0ff_SmartRecruiters.png"
              alt="brand-logo"
            />
            <img
              src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f4e5b56269716862f12fd40_slido_black.png"
              alt="brand-logo"
            />
          </s.BrandImages>
        </s.LearnFromBestLogos>
        <s.SectionFutureSummit>
          <s.ContentLeftWrapper>
            <s.ContentSubTitle>Remote Future Summit</s.ContentSubTitle>
            <h3>Not just another virtual conference.</h3>
            <p>
              We live in a time filled with seemingly hundreds of digital
              products and tools to solve issues for remote workers. That’s why
              we want Remote Future Summit 2020 to be an event to connect with
              each other, be sincere about struggles, and discuss our
              relationships with remote teams, decreased visibility, and
              wellbeing.
            </p>
            <button>Register Now</button>
          </s.ContentLeftWrapper>
          <s.VideoWrapper>
            <h1>Video Here</h1>
          </s.VideoWrapper>
        </s.SectionFutureSummit>
        <s.SectionLargeText>
          <s.TextBlack>
            Over the past years different studies have always pointed to these
            struggles among remote workers:{" "}
          </s.TextBlack>
          <s.TextBlue>
            overworking, difficulties with building relationships, and with
            loneliness.
          </s.TextBlue>
          <s.TextGrey>We will help you overcome these challenges.</s.TextGrey>
        </s.SectionLargeText>
        <s.SectionWhatYouWillGet>
          <s.ContentSubTitle>Remote Future Summit</s.ContentSubTitle>
          <h3>What you will get?</h3>
          <s.Boxes>
            <s.HundredPercentLive>
              <s.CircleImageWrapper>
                <img
                  width="44"
                  src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f329fb1fd8a91959a5cd40b_smartphone-tablet.svg"
                />
              </s.CircleImageWrapper>
              <h4>100% live</h4>
              <p>
                Live-streamed and interactive sessions for two days in a row!
              </p>
            </s.HundredPercentLive>
            <s.GoodVibes>
              <s.CircleImageWrapper>
                <img
                  width="44"
                  src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f329fb1fd8a9157e55cd435_icon-heart.svg"
                />
              </s.CircleImageWrapper>
              <h4>Good vibes only</h4>
              <p>Yoga and coffee breaks where we recharge our batteries</p>
            </s.GoodVibes>
            <s.NoBullShit>
              <s.CircleImageWrapper>
                <img
                  width="44"
                  src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f329fb1fd8a9157ca5cd433_icon-x.svg"
                />
              </s.CircleImageWrapper>
              <h4>No bulsh*t</h4>
              <p>
                Industry experts sharing openly about their challenges &
                failures
              </p>
            </s.NoBullShit>
          </s.Boxes>
        </s.SectionWhatYouWillGet>
        <s.SectionSpeakersContainer>
          <s.ContentSubTitle>They will guide you</s.ContentSubTitle>
          <h3>Meet the Speakers</h3>
          <s.SpeakersWrapper>
            <s.Speaker>
              <img
                src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                alt="speaker-image"
              />
              <h4>Jerome Ternynck</h4>
              <s.TextLead>
                Founder &amp; CEO <br />
                Smart Recruiters
              </s.TextLead>
              <s.SpeakerSocialMedia>
                <a href={"https://twitter.com/elonmusk"}>
                  <i className="fa fa-twitter twitter" aria-hidden="true"></i>
                </a>
                <a href={"https://linkedin.com"}>
                  <i className="fa fa-linkedin linkedin" aria-hidden="true"></i>
                </a>
              </s.SpeakerSocialMedia>
            </s.Speaker>
            <s.Speaker>
              <img
                src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                alt="speaker-image"
              />
              <h4>Jerome Ternynck</h4>
              <s.TextLead>
                Founder &amp; CEO <br />
                Smart Recruiters
              </s.TextLead>
              <s.SpeakerSocialMedia>
                <a href={"https://twitter.com/elonmusk"}>
                  <i className="fa fa-twitter twitter" aria-hidden="true"></i>
                </a>
                <a href={"https://linkedin.com"}>
                  <i className="fa fa-linkedin linkedin" aria-hidden="true"></i>
                </a>
              </s.SpeakerSocialMedia>
            </s.Speaker>
            <s.Speaker>
              <img
                src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                alt="speaker-image"
              />
              <h4>Jerome Ternynck</h4>
              <s.TextLead>
                Founder &amp; CEO <br />
                Smart Recruiters
              </s.TextLead>
              <s.SpeakerSocialMedia>
                <a href={"https://twitter.com/elonmusk"}>
                  <i className="fa fa-twitter twitter" aria-hidden="true"></i>
                </a>
                <a href={"https://linkedin.com"}>
                  <i className="fa fa-linkedin linkedin" aria-hidden="true"></i>
                </a>
              </s.SpeakerSocialMedia>
            </s.Speaker>
            <s.Speaker>
              <img
                src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                alt="speaker-image"
              />
              <h4>Jerome Ternynck</h4>
              <s.TextLead>
                Founder &amp; CEO <br />
                Smart Recruiters
              </s.TextLead>
              <s.SpeakerSocialMedia>
                <a href={"https://twitter.com/elonmusk"}>
                  <i className="fa fa-twitter twitter" aria-hidden="true"></i>
                </a>
                <a href={"https://linkedin.com"}>
                  <i className="fa fa-linkedin linkedin" aria-hidden="true"></i>
                </a>
              </s.SpeakerSocialMedia>
            </s.Speaker>
            <s.Speaker>
              <img
                src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                alt="speaker-image"
              />
              <h4>Jerome Ternynck</h4>
              <s.TextLead>
                Founder &amp; CEO <br />
                Smart Recruiters
              </s.TextLead>
              <s.SpeakerSocialMedia>
                <a href={"https://twitter.com/elonmusk"}>
                  <i className="fa fa-twitter twitter" aria-hidden="true"></i>
                </a>
                <a href={"https://linkedin.com"}>
                  <i className="fa fa-linkedin linkedin" aria-hidden="true"></i>
                </a>
              </s.SpeakerSocialMedia>
            </s.Speaker>
            <s.Speaker>
              <img
                src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                alt="speaker-image"
              />
              <h4>Jerome Ternynck</h4>
              <s.TextLead>
                Founder &amp; CEO <br />
                Smart Recruiters
              </s.TextLead>
              <s.SpeakerSocialMedia>
                <a href={"https://twitter.com/elonmusk"}>
                  <i className="fa fa-twitter twitter" aria-hidden="true"></i>
                </a>
                <a href={"https://linkedin.com"}>
                  <i className="fa fa-linkedin linkedin" aria-hidden="true"></i>
                </a>
              </s.SpeakerSocialMedia>
            </s.Speaker>
            <s.Speaker>
              <img
                src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                alt="speaker-image"
              />
              <h4>Jerome Ternynck</h4>
              <s.TextLead>
                Founder &amp; CEO <br />
                Smart Recruiters
              </s.TextLead>
              <s.SpeakerSocialMedia>
                <a href={"https://twitter.com/elonmusk"}>
                  <i className="fa fa-twitter twitter" aria-hidden="true"></i>
                </a>
                <a href={"https://linkedin.com"}>
                  <i className="fa fa-linkedin linkedin" aria-hidden="true"></i>
                </a>
              </s.SpeakerSocialMedia>
            </s.Speaker>
            <s.Speaker>
              <img
                src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                alt="speaker-image"
              />
              <h4>Jerome Ternynck</h4>
              <s.TextLead>
                Founder &amp; CEO <br />
                Smart Recruiters
              </s.TextLead>
              <s.SpeakerSocialMedia>
                <a href={"https://twitter.com/elonmusk"}>
                  <i className="fa fa-twitter twitter" aria-hidden="true"></i>
                </a>
                <a href={"https://linkedin.com"}>
                  <i className="fa fa-linkedin linkedin" aria-hidden="true"></i>
                </a>
              </s.SpeakerSocialMedia>
            </s.Speaker>
          </s.SpeakersWrapper>
        </s.SectionSpeakersContainer>
        <s.SectionAgenda>
          <h1>See The Agenda</h1>
          <s.TabsWrapper>
            <Tabs>
              <TabList>
                <Tab>Day 1</Tab>
                <Tab>Day 2</Tab>
              </TabList>
              <s.TabContent>
                <TabPanel>
                  <s.EventTitle>
                    <h3>Day 1</h3>
                  </s.EventTitle>
                  <s.EventTime>
                    <h2>DAY 1 Opening - 9:00 AM</h2>
                  </s.EventTime>
                  <s.PanelDiscussion>
                    <s.topic>
                      Panel discussion: What pandemic taught us about remote
                      work
                      <s.PanelImages>
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                      </s.PanelImages>
                    </s.topic>
                    <s.TimeLimit>9:10 AM-9:50 AM</s.TimeLimit>
                  </s.PanelDiscussion>
                  <s.PanelDiscussion>
                    <s.topic>
                      Panel discussion: What pandemic taught us about remote
                      work
                      <s.PanelImages>
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                      </s.PanelImages>
                    </s.topic>
                    <s.TimeLimit>9:10 AM-9:50 AM</s.TimeLimit>
                  </s.PanelDiscussion>
                  <s.PanelDiscussion>
                    <s.topic>
                      Panel discussion: What pandemic taught us about remote
                      work
                      <s.PanelImages>
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                      </s.PanelImages>
                    </s.topic>
                    <s.TimeLimit>9:10 AM-9:50 AM</s.TimeLimit>
                  </s.PanelDiscussion>
                </TabPanel>
                <TabPanel>
                  <s.EventTitle>
                    <h3>Day 2</h3>
                  </s.EventTitle>
                  <s.EventTime>
                    <h2>DAY 2 Opening - 9:00 AM</h2>
                  </s.EventTime>
                  <s.PanelDiscussion>
                    <s.topic>
                      Panel discussion: What pandemic taught us about remote
                      work
                      <s.PanelImages>
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                      </s.PanelImages>
                    </s.topic>
                    <s.TimeLimit>9:10 AM-9:50 AM</s.TimeLimit>
                  </s.PanelDiscussion>
                  <s.PanelDiscussion>
                    <s.topic>
                      Panel discussion: What pandemic taught us about remote
                      work
                      <s.PanelImages>
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                      </s.PanelImages>
                    </s.topic>
                    <s.TimeLimit>9:10 AM-9:50 AM</s.TimeLimit>
                  </s.PanelDiscussion>
                  <s.PanelDiscussion>
                    <s.topic>
                      Panel discussion: What pandemic taught us about remote
                      work
                      <s.PanelImages>
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                        <img
                          src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
                          alt="speaker-image"
                        />
                      </s.PanelImages>
                    </s.topic>
                    <s.TimeLimit>9:10 AM-9:50 AM</s.TimeLimit>
                  </s.PanelDiscussion>
                </TabPanel>
                <s.KeepInMindText>
                  "❗ Keep in mind all session timings are displayed in 🕤
                  Eastern Time Zone (ET) 🕛 ❗"{" "}
                </s.KeepInMindText>
              </s.TabContent>
            </Tabs>
          </s.TabsWrapper>
        </s.SectionAgenda>
        <s.SectionIsThisForMe>
          <s.ContentSubTitle>Is this for me?</s.ContentSubTitle>
          <h3>Remote Future Summit 2020 is made for</h3>
        </s.SectionIsThisForMe>
        <s.ContentIsThisForMe>
          <s.FounderAndCeo>
            <h4>Founders & CEOs of remote companies</h4>
            <s.Badges>
              <s.Badge>Day 1</s.Badge>
              <s.Badge>Day 2</s.Badge>
            </s.Badges>
            <p>
              Learn valuable insights and get answers on the reality of working
              remotely and leading virtual teams directly from invited leaders
              and seasoned executives.{" "}
            </p>
          </s.FounderAndCeo>
          <s.FounderAndCeo>
            <h4>Founders & CEOs of remote companies</h4>
            <s.Badges>
              <s.Badge>Day 1</s.Badge>
              <s.Badge>Day 2</s.Badge>
            </s.Badges>
            <p>
              Learn valuable insights and get answers on the reality of working
              remotely and leading virtual teams directly from invited leaders
              and seasoned executives.{" "}
            </p>
          </s.FounderAndCeo>
          <s.FounderAndCeo>
            <h4>Founders & CEOs of remote companies</h4>
            <s.Badges>
              <s.Badge>Day 1</s.Badge>
              <s.Badge>Day 2</s.Badge>
            </s.Badges>
            <p>
              Learn valuable insights and get answers on the reality of working
              remotely and leading virtual teams directly from invited leaders
              and seasoned executives.{" "}
            </p>
          </s.FounderAndCeo>
        </s.ContentIsThisForMe>
        <s.ButtonSaveYourSeatNow>
          <button>Save your seat now</button>
        </s.ButtonSaveYourSeatNow>
        <s.SectionAboutOrganizer>
          <s.AboutOrganizer>
            <s.ContentSubTitle>About organizers</s.ContentSubTitle>
            <h3>We've been in this since 2017'</h3>
            <p>
              This will be the third edition of our virtual conference on remote
              work powered by Remote-how. In 2018 and 2019 we had over 8,000
              attendees from 105 countries, with 64 speakers including Asana,
              Hubspot, Forbes, and Doist. Our goal is to help you learn how to
              build, scale, and lead distributed teams of remote workers.
            </p>
          </s.AboutOrganizer>
          <s.OrganizerImage>
            {" "}
            <img
              src="https://uploads-ssl.webflow.com/5f329fb0017255d9d0baddec/5f3a7e64ecda612e4c4ab82e_Jerome_Remote%20Future%20Summit.jpg"
              alt="speaker-image"
            />
          </s.OrganizerImage>
        </s.SectionAboutOrganizer>
        <s.Partners>
          <h3>Partners</h3>
          <p> We will show some partner logos here</p>
        </s.Partners>
        <s.FrequentlyQA>
          <h3>Frequently asked questions</h3>
          <s.Question>Where and when is Remote Future Summit 2020?</s.Question>
          <s.Answer>
            Remote Future Summit 2020 is a virtual conference. It’s happening
            online from September 23-24, 2020. You can join it from anywhere,
            literally.{" "}
          </s.Answer>
        </s.FrequentlyQA>
        <s.Footer>
          <h3>Contact Us</h3>
          <p>support@campsite.org</p>
          <s.SpeakerSocialMedia>
            <a href={"https://twitter.com/elonmusk"}>
              <i className="fa fa-twitter twitter" aria-hidden="true"></i>
            </a>
            <a href={"https://linkedin.com"}>
              <i className="fa fa-linkedin linkedin" aria-hidden="true"></i>
            </a>
            <a href={"https://instagram.com"}>
              <i className="fa fa-instagram instagram" aria-hidden="true"></i>
            </a>
          </s.SpeakerSocialMedia>

          <s.PoweredBy>
            Powered by <a href={"https://campsite.com"}>Campsite</a>
            <div className="privacy-policy">
              <a href={"https://campsite.org/terms"}>Terms of Use</a>
              <a href={"https://campsite.org/privacy"}>Privacy Policy</a>
            </div>
          </s.PoweredBy>
        </s.Footer>
      </g.Container>
    </Fragment>
  );
}

export default HomePage;
