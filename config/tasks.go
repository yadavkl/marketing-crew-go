package config

import "time"

type Input struct {
	ProductName        string
	TargetAudience     string
	ProductDescription string
	Budget             string
	CurrentDate        time.Time
	Other              string
}

type Output struct {
	Description     string
	Expectectations []string
	Comment         string
}
type Task struct {
	Description    string
	Input          Input
	ExpectedOutput Output
	Comment        string
}

func New() *Task {
	return &Task{}
}

func (t *Task) WithDescription(description string) *Task {
	t.Description = description
	return t
}
func (t *Task) WithInput(input Input) *Task {
	t.Input = input
	return t
}

func (t *Task) WithExpectedOutput(output Output) *Task {
	t.ExpectedOutput = output
	return t
}

func (t *Task) WithComment(comment string) *Task {
	t.Comment = comment
	return t
}

var MarkteResearchTask = New().
	WithDescription("Conduct market research to identify trends, opportunities, and challenges in the industry. Focus on understanding customer needs, competitor strategies, and market dynamics to inform marketing decisions.").
	WithExpectedOutput(Output{
		Description: "A comprehensive market research report that includes:",
		Expectectations: []string{
			"Analysis of market trends and opportunities",
			"Competitor analysis",
			"Customer insights and needs",
			"Recommendations for marketing strategies",
		},
	})

var PrepareMarketingStrategyTask = New().
	WithDescription("Develop a marketing strategy that aligns with the company's goals and objectives. The strategy should include target audience segmentation, positioning, messaging, and marketing channels.").
	WithExpectedOutput(Output{
		Description: "A detailed marketing strategy document that includes:",
		Expectectations: []string{
			"Target audience segmentation",
			"Required budget",
			"Recommended marketing channels",
			"Plan of action for the week containing 3 blogs to write, 2 reels to film, and 5 social media posts to create",
			"Key performance indicators (KPIs) to measure success",
		},
		Comment: "Store the marketing strategy in 'resources/drafts/marketing_strategy.md'.",
	}).
	WithComment("Refer to the market research report for insights and building the strategy.")

var CreateContentCalenderTask = New().
	WithDescription("Create a content calendar that outlines the topics, formats, and publishing schedule for the next one week. The calendar should align with the marketing strategy and include key themes and campaigns.").
	WithExpectedOutput(Output{
		Description: " A content calendar that includes:",
		Expectectations: []string{
			"Topics and formats for each piece of content specified in the marketing strategy",
			"Publishing schedule",
			"Key themes and campaigns",
		},
		Comment: "Store the content calendar in 'resources/drafts/content_calendar.md'.",
	})

var PreparePostDraftTask = New().
	WithDescription("Prepare drafts for social media posts and email campaigns based on the content calendar. Ensure that the drafts are engaging, aligned with the brand voice, and optimized for each platform.").
	WithExpectedOutput(Output{
		Description: "Drafts for:",
		Expectectations: []string{
			"Social media posts: LinkedIn, Twitter, Instagram",
			"Email campaigns",
		},
		Comment: "Each draft should be ready for review and feedback. Save the drafts in 'resources/drafts/posts/' folder in markdown format.",
	})

var PrepareScriptsForReelsTask = New().
	WithDescription("Prepare scripts for Instagram reels that align with the content calendar and marketing strategy. The scripts should be engaging, concise, and tailored to the target audience.").
	WithExpectedOutput(Output{
		Description: " Scripts for Instagram reels that include:",
		Expectectations: []string{
			"Engaging hooks",
			"Key messages",
			"Call to action",
		},
		Comment: "Each script should be ready for filming and editing. Save the scripts in 'resources/drafts/reels/' folder in markdown format. And assign tasks to the relevant team members.",
	})

var ContentResearchForBlogsTask = New().
	WithDescription("Conduct research for blog topics that align with the content calendar and marketing strategy. The research should include keyword analysis, competitor blogs, and industry trends.").
	WithExpectedOutput(Output{
		Description: " A research report that includes:",
		Expectectations: []string{
			"Keyword analysis",
			"Competitor blog insights",
			"Industry trends",
			"Suggested blog topics and outlines",
		},
	})
var DraftBlogsTask = New().
	WithDescription("Draft blogs based on the research report and content calendar. Each blog should be well-structured, engaging, and optimized for SEO.").
	WithExpectedOutput(Output{
		Description: "Drafts for blogs that include:",
		Expectectations: []string{
			"Engaging introductions",
			"Well-structured content",
			"SEO optimization",
			"Ready for review and feedback",
		},
		Comment: "Each draft should be ready for review and saved in 'resources/drafts/blogs/' folder in markdown format.",
	})

var SeoOptimizationTask = New().
	WithDescription("Optimize the drafted blogs for search engines to improve visibility and drive organic traffic. This includes keyword optimization, meta tags, and internal linking.").
	WithExpectedOutput(Output{
		Description: "SEO-optimized blogs that include:",
		Expectectations: []string{
			"Keyword-rich titles and headings",
			"Meta descriptions",
			"Internal links",
			"Ready for final review and publication",
		},
	})
