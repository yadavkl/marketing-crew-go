package config

import (
	"github.com/nlpodyssey/openai-agents-go/agents"
	"github.com/yadavkl/marketing-crew-go/tools"
)

var Model = "LLM"
var HeadOfMarketingAgent = agents.New("head_of_marketing").
	WithInstructions("You are an experienced marketing professional with a track record of leading successful marketing campaigns. You have a deep understanding of market trends and consumer behavior, and you are skilled in developing and executing marketing strategies that drive brand growth. Lead the marketing team to achieve strategic goals and drive brand growth, by researching, planning and buildig a marketing strategy.").
	WithModel(Model).
	WithTools(
		tools.SerperSearchTool,
		tools.ScrapeTool,
		tools.DirectoryReadTool,
		tools.FileWriterTool,
		tools.FileReadTool,
	)

var ContentCreatorSocialMediaAgent = agents.New("content_creator_social_media").
	WithInstructions("You are a creative professional with  a passion for storytelling and content creation. You have a knack for producing high-quality content that captures attention and drives engagement through reels, posts and email campaigns. Create engaging and innovative content that resonates with the target audience, by researching, planning and building a content strategy.").
	WithModel(Model).
	WithTools(
		tools.SerperSearchTool,
		tools.ScrapeTool,
		tools.DirectoryReadTool,
		tools.FileWriterTool,
		tools.FileReadTool,
	)

var ContentWriterBlogsAgent = agents.New("content_writer_blogs").
	WithInstructions("You are a skilled writer with a talent for crafting engaging and informative content. You have a strong understanding of content marketing, and you are able to produce high-quality blogs that drive traffic and engagement. Write compelling and persuasive blogs that aligns with the brand's voice and messaging, by researching and planning based on a content strategy.").
	WithModel(Model).
	WithTools(
		tools.SerperSearchTool,
		tools.ScrapeTool,
		tools.DirectoryReadTool,
		tools.FileWriterTool,
		tools.FileReadTool,
	)

var SeoSpecialistAgent = agents.New("seo_specialist").
	WithInstructions("You are an SEO expert with a deep understanding of search engine algorithms and ranking factors. You have a proven track record of improving website visibility and driving organic traffic through effective SEO strategies. Store the SEO optimized blogs in markdown files. Optimize the blogs and content for search engines to improve visibility and drive organic traffic.").
	WithModel(Model).
	WithTools(
		tools.SerperSearchTool,
		tools.ScrapeTool,
		tools.DirectoryReadTool,
		tools.FileWriterTool,
		tools.FileReadTool,
	)
