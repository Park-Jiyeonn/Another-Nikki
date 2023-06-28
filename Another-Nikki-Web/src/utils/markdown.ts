import MarkdownIt from "markdown-it";
import hljs from 'highlight.js/lib/common'
import markdownItAttrs from 'markdown-it-attrs'

import mdKatex from 'markdown-it-katex'

import 'highlight.js/styles/atom-one-dark.css'

export const md = new MarkdownIt({
	html: true, // 在源码中启用HTML标签
	linkify: true, // 将类似URL的文本自动转换为链接
	breaks: true, // 转换段落里的 '\n' 到 <br>
	highlight: function (str, lang) {
		return highlightFormatCode(str, lang)
	}
}) as MarkdownIt
md.use(mdKatex, {"throwOnError" : false, "errorColor" : " #cc0000"})
md.use(markdownItAttrs);

const highlightFormatCode = (str: string, lang: string): string => {
	if (lang && hljs.getLanguage(lang)) {
		try {
			return codeBlockStyle(hljs.highlight(str, { language: lang, ignoreIllegals: true }).value)
		} catch (e) {
			console.error(e)
		}
	}

	return codeBlockStyle(md.utils.escapeHtml(str))
}

const codeBlockStyle = (val: string): string => {
	return `<pre class="hljs" style="padding: 10px;border-radius: 10px;"><code>${val}</code></pre>`
}


export const detectMarkdown = (text: string) => {
	const reHeading = /^\s*#+\s*.+$/mg
	const reCodeBlockFence = /^\s*`{3,}\w*\s*$/mg
	const reImage = /<\s*img\s+.+>/
	const reBr = /<\s*br\s*\/?\s*>/g

	if (_matchCount(text, reHeading) > 1) {
		return true
	}
	else if (_matchCount(text, reCodeBlockFence) > 1) {
		return true
	}
	else if (_matchCount(text, reBr) > 1) {
		return true
	}
	else if (reImage.test(text)) {
		return true
	}
	return false
}

function _matchCount(str: string, regex: RegExp) {
	const matches = str.match(regex) || []
	return matches.length
}


export const renderMarkdown = (markdown: string) => {
	let ret = md.render(markdown)
	// ret = htmlDecode(ret)
	return ret
}

// export const htmlDecode = (text : string) : string => {
// 	let temp = document.createElement('div')
// 	temp.innerHTML = text;
// 	let output = temp.innerText || temp.textContent;

// 	return output;
// }
