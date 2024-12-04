import{_ as a,c as n,a0 as l,o as p}from"./chunks/framework.iBbVlLT2.js";const u=JSON.parse('{"title":"Use in Go","description":"","frontmatter":{},"headers":[],"relativePath":"introduction/go-module.md","filePath":"introduction/go-module.md"}'),o={name:"introduction/go-module.md"};function t(e,s,F,c,r,y){return p(),n("div",null,s[0]||(s[0]=[l(`<h1 id="use-in-go" tabindex="-1">Use in Go <a class="header-anchor" href="#use-in-go" aria-label="Permalink to &quot;Use in Go&quot;">​</a></h1><p>Mugo&#39;s internal template runner and functions are available as a Go package.</p><h2 id="templatex" tabindex="-1">templatex <a class="header-anchor" href="#templatex" aria-label="Permalink to &quot;templatex&quot;">​</a></h2><p><code>templatex</code> is a package that provides a template runner with options.</p><div class="language-sh"><button title="Copy Code" class="copy"></button><span class="lang">sh</span><pre class="shiki dracula vp-code" tabindex="0"><code><span class="line"><span style="color:#50FA7B;">go</span><span style="color:#F1FA8C;"> get</span><span style="color:#F1FA8C;"> github.com/rytsh/mugo</span></span></code></pre></div><div class="language-sh"><button title="Copy Code" class="copy"></button><span class="lang">sh</span><pre class="shiki dracula vp-code" tabindex="0"><code><span class="line"><span style="color:#50FA7B;">import</span><span style="color:#E9F284;"> &quot;</span><span style="color:#F1FA8C;">github.com/rytsh/mugo/templatex</span><span style="color:#E9F284;">&quot;</span></span></code></pre></div><h3 id="usage" tabindex="-1">Usage <a class="header-anchor" href="#usage" aria-label="Permalink to &quot;Usage&quot;">​</a></h3><p>Check details in go document: <a href="https://pkg.go.dev/github.com/rytsh/mugo/templatex" target="_blank" rel="noreferrer">https://pkg.go.dev/github.com/rytsh/mugo/templatex</a></p><div class="language-go"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki dracula vp-code" tabindex="0"><code><span class="line"><span style="color:#F8F8F2;">tpl </span><span style="color:#FF79C6;">:=</span><span style="color:#F8F8F2;"> templatex.</span><span style="color:#50FA7B;">New</span><span style="color:#F8F8F2;">()</span></span>
<span class="line"></span>
<span class="line"><span style="color:#F8F8F2;">tpl.</span><span style="color:#50FA7B;">AddFunc</span><span style="color:#F8F8F2;">(</span><span style="color:#E9F284;">&quot;</span><span style="color:#F1FA8C;">add</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">, </span><span style="color:#FF79C6;">func</span><span style="color:#F8F8F2;">(</span><span style="color:#FFB86C;font-style:italic;">a</span><span style="color:#F8F8F2;">, </span><span style="color:#FFB86C;font-style:italic;">b</span><span style="color:#8BE9FD;font-style:italic;"> int</span><span style="color:#F8F8F2;">) </span><span style="color:#8BE9FD;font-style:italic;">int</span><span style="color:#F8F8F2;"> {</span></span>
<span class="line"><span style="color:#FF79C6;">    return</span><span style="color:#F8F8F2;"> a </span><span style="color:#FF79C6;">+</span><span style="color:#F8F8F2;"> b</span></span>
<span class="line"><span style="color:#F8F8F2;">})</span></span>
<span class="line"></span>
<span class="line"><span style="color:#F8F8F2;">tpl.</span><span style="color:#50FA7B;">AddFunc</span><span style="color:#F8F8F2;">(</span><span style="color:#E9F284;">&quot;</span><span style="color:#F1FA8C;">sub</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">, </span><span style="color:#FF79C6;">func</span><span style="color:#F8F8F2;">(</span><span style="color:#FFB86C;font-style:italic;">a</span><span style="color:#F8F8F2;">, </span><span style="color:#FFB86C;font-style:italic;">b</span><span style="color:#8BE9FD;font-style:italic;"> int</span><span style="color:#F8F8F2;">) </span><span style="color:#8BE9FD;font-style:italic;">int</span><span style="color:#F8F8F2;"> {</span></span>
<span class="line"><span style="color:#FF79C6;">    return</span><span style="color:#F8F8F2;"> a </span><span style="color:#FF79C6;">-</span><span style="color:#F8F8F2;"> b</span></span>
<span class="line"><span style="color:#F8F8F2;">})</span></span>
<span class="line"></span>
<span class="line"><span style="color:#FF79C6;">var</span><span style="color:#F8F8F2;"> output </span><span style="color:#8BE9FD;font-style:italic;">bytes</span><span style="color:#F8F8F2;">.</span><span style="color:#8BE9FD;font-style:italic;">Buffer</span></span>
<span class="line"><span style="color:#FF79C6;">if</span><span style="color:#F8F8F2;"> err </span><span style="color:#FF79C6;">:=</span><span style="color:#F8F8F2;"> tpl.</span><span style="color:#50FA7B;">Execute</span><span style="color:#F8F8F2;">(</span></span>
<span class="line"><span style="color:#F8F8F2;">    templatex.</span><span style="color:#50FA7B;">WithIO</span><span style="color:#F8F8F2;">(</span><span style="color:#FF79C6;">&amp;</span><span style="color:#F8F8F2;">output),</span></span>
<span class="line"><span style="color:#F8F8F2;">    templatex.</span><span style="color:#50FA7B;">WithData</span><span style="color:#F8F8F2;">(</span><span style="color:#FF79C6;">map</span><span style="color:#F8F8F2;">[</span><span style="color:#8BE9FD;font-style:italic;">string</span><span style="color:#F8F8F2;">]</span><span style="color:#FF79C6;">interface</span><span style="color:#F8F8F2;">{}{</span></span>
<span class="line"><span style="color:#E9F284;">        &quot;</span><span style="color:#F1FA8C;">a</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">: </span><span style="color:#BD93F9;">1</span><span style="color:#F8F8F2;">,</span></span>
<span class="line"><span style="color:#E9F284;">        &quot;</span><span style="color:#F1FA8C;">b</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">: </span><span style="color:#BD93F9;">2</span><span style="color:#F8F8F2;">,</span></span>
<span class="line"><span style="color:#F8F8F2;">    }),</span></span>
<span class="line"><span style="color:#F8F8F2;">    templatex.</span><span style="color:#50FA7B;">WithContent</span><span style="color:#F8F8F2;">(</span><span style="color:#E9F284;">\`</span><span style="color:#F1FA8C;">a + b = {{ add .a .b }}</span><span style="color:#E9F284;">\`</span><span style="color:#FF79C6;">+</span><span style="color:#E9F284;">&quot;</span><span style="color:#FF79C6;">\\n</span><span style="color:#E9F284;">&quot;</span><span style="color:#FF79C6;">+</span><span style="color:#E9F284;">\`</span><span style="color:#F1FA8C;">a - b = {{ sub .a .b }}</span><span style="color:#E9F284;">\`</span><span style="color:#F8F8F2;">),</span></span>
<span class="line"><span style="color:#F8F8F2;">); err </span><span style="color:#FF79C6;">!=</span><span style="color:#BD93F9;"> nil</span><span style="color:#F8F8F2;"> {</span></span>
<span class="line"><span style="color:#F8F8F2;">    log.</span><span style="color:#50FA7B;">Fatal</span><span style="color:#F8F8F2;">(err)</span></span>
<span class="line"><span style="color:#F8F8F2;">}</span></span>
<span class="line"></span>
<span class="line"><span style="color:#F8F8F2;">fmt.</span><span style="color:#50FA7B;">Printf</span><span style="color:#F8F8F2;">(</span><span style="color:#E9F284;">&quot;</span><span style="color:#BD93F9;">%s</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">, output.</span><span style="color:#50FA7B;">String</span><span style="color:#F8F8F2;">())</span></span>
<span class="line"><span style="color:#6272A4;">// Output:</span></span>
<span class="line"><span style="color:#6272A4;">// a + b = 3</span></span>
<span class="line"><span style="color:#6272A4;">// a - b = -1</span></span></code></pre></div><h2 id="fstore" tabindex="-1">fstore <a class="header-anchor" href="#fstore" aria-label="Permalink to &quot;fstore&quot;">​</a></h2><p><code>fstore</code> is a package that provides bunch of functions with options.</p><div class="language-sh"><button title="Copy Code" class="copy"></button><span class="lang">sh</span><pre class="shiki dracula vp-code" tabindex="0"><code><span class="line"><span style="color:#50FA7B;">go</span><span style="color:#F1FA8C;"> get</span><span style="color:#F1FA8C;"> github.com/rytsh/mugo</span></span></code></pre></div><div class="language-sh"><button title="Copy Code" class="copy"></button><span class="lang">sh</span><pre class="shiki dracula vp-code" tabindex="0"><code><span class="line"><span style="color:#50FA7B;">import</span><span style="color:#E9F284;"> &quot;</span><span style="color:#F1FA8C;">github.com/rytsh/mugo/fstore</span><span style="color:#E9F284;">&quot;</span></span></code></pre></div><h3 id="usage-1" tabindex="-1">Usage <a class="header-anchor" href="#usage-1" aria-label="Permalink to &quot;Usage&quot;">​</a></h3><p>Check details in go document: <a href="https://pkg.go.dev/github.com/rytsh/mugo/fstore" target="_blank" rel="noreferrer">https://pkg.go.dev/github.com/rytsh/mugo/fstore</a></p><div class="language-go"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki dracula vp-code" tabindex="0"><code><span class="line"><span style="color:#F8F8F2;">tpl </span><span style="color:#FF79C6;">:=</span><span style="color:#F8F8F2;"> template.</span><span style="color:#50FA7B;">New</span><span style="color:#F8F8F2;">(</span><span style="color:#E9F284;">&quot;</span><span style="color:#F1FA8C;">test</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">).</span><span style="color:#50FA7B;">Funcs</span><span style="color:#F8F8F2;">(fstore.</span><span style="color:#50FA7B;">FuncMap</span><span style="color:#F8F8F2;">())</span></span>
<span class="line"></span>
<span class="line"><span style="color:#F8F8F2;">output </span><span style="color:#FF79C6;">:=</span><span style="color:#FF79C6;"> &amp;</span><span style="color:#8BE9FD;font-style:italic;">bytes</span><span style="color:#F8F8F2;">.</span><span style="color:#8BE9FD;font-style:italic;">Buffer</span><span style="color:#F8F8F2;">{}</span></span>
<span class="line"><span style="color:#F8F8F2;">tplParsed, err </span><span style="color:#FF79C6;">:=</span><span style="color:#F8F8F2;"> tpl.</span><span style="color:#50FA7B;">Parse</span><span style="color:#F8F8F2;">(</span><span style="color:#E9F284;">\`</span><span style="color:#F1FA8C;">{{ $v := codec.JsonDecode (codec.StringToByte .) }}{{ $v.data.name }}</span><span style="color:#E9F284;">\`</span><span style="color:#F8F8F2;">)</span></span>
<span class="line"><span style="color:#FF79C6;">if</span><span style="color:#F8F8F2;"> err </span><span style="color:#FF79C6;">!=</span><span style="color:#BD93F9;"> nil</span><span style="color:#F8F8F2;"> {</span></span>
<span class="line"><span style="color:#F8F8F2;">    log.</span><span style="color:#50FA7B;">Fatal</span><span style="color:#F8F8F2;">(err)</span></span>
<span class="line"><span style="color:#F8F8F2;">}</span></span>
<span class="line"></span>
<span class="line"><span style="color:#FF79C6;">if</span><span style="color:#F8F8F2;"> err </span><span style="color:#FF79C6;">:=</span><span style="color:#F8F8F2;"> tplParsed.</span><span style="color:#50FA7B;">Execute</span><span style="color:#F8F8F2;">(output, </span><span style="color:#E9F284;">\`</span><span style="color:#F1FA8C;">{&quot;data&quot;: {&quot;name&quot;: &quot;Hatay&quot;}}</span><span style="color:#E9F284;">\`</span><span style="color:#F8F8F2;">); err </span><span style="color:#FF79C6;">!=</span><span style="color:#BD93F9;"> nil</span><span style="color:#F8F8F2;"> {</span></span>
<span class="line"><span style="color:#F8F8F2;">    log.</span><span style="color:#50FA7B;">Fatal</span><span style="color:#F8F8F2;">(err)</span></span>
<span class="line"><span style="color:#F8F8F2;">}</span></span>
<span class="line"></span>
<span class="line"><span style="color:#F8F8F2;">fmt.</span><span style="color:#50FA7B;">Printf</span><span style="color:#F8F8F2;">(</span><span style="color:#E9F284;">&quot;</span><span style="color:#BD93F9;">%s</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">, output)</span></span>
<span class="line"><span style="color:#6272A4;">// Output:</span></span>
<span class="line"><span style="color:#6272A4;">// Hatay</span></span></code></pre></div><h1 id="use-fstore-with-templatex" tabindex="-1">Use fstore with templatex <a class="header-anchor" href="#use-fstore-with-templatex" aria-label="Permalink to &quot;Use fstore with templatex&quot;">​</a></h1><p><code>fstore</code> and <code>templatex</code> can be used together. Use the tpl to execute templates.</p><div class="language-go"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki dracula vp-code" tabindex="0"><code><span class="line"><span style="color:#F8F8F2;">tpl </span><span style="color:#FF79C6;">:=</span><span style="color:#F8F8F2;"> templatex.</span><span style="color:#50FA7B;">New</span><span style="color:#F8F8F2;">(templatex.</span><span style="color:#50FA7B;">WithAddFuncsTpl</span><span style="color:#F8F8F2;">(</span></span>
<span class="line"><span style="color:#F8F8F2;">    fstore.</span><span style="color:#50FA7B;">FuncMapTpl</span><span style="color:#F8F8F2;">(</span></span>
<span class="line"><span style="color:#F8F8F2;">        fstore.</span><span style="color:#50FA7B;">WithLog</span><span style="color:#F8F8F2;">(</span><span style="color:#8BE9FD;font-style:italic;">logz</span><span style="color:#F8F8F2;">.</span><span style="color:#8BE9FD;font-style:italic;">AdapterKV</span><span style="color:#F8F8F2;">{Log: log.Logger}),</span></span>
<span class="line"><span style="color:#F8F8F2;">        fstore.</span><span style="color:#50FA7B;">WithTrust</span><span style="color:#F8F8F2;">(</span><span style="color:#BD93F9;">true</span><span style="color:#F8F8F2;">),</span></span>
<span class="line"><span style="color:#F8F8F2;">        fstore.</span><span style="color:#50FA7B;">WithWorkDir</span><span style="color:#F8F8F2;">(</span><span style="color:#E9F284;">&quot;</span><span style="color:#F1FA8C;">.</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">),</span></span>
<span class="line"><span style="color:#F8F8F2;">    ),</span></span>
<span class="line"><span style="color:#F8F8F2;">))</span></span>
<span class="line"></span>
<span class="line"><span style="color:#FF79C6;">var</span><span style="color:#F8F8F2;"> buf </span><span style="color:#8BE9FD;font-style:italic;">bytes</span><span style="color:#F8F8F2;">.</span><span style="color:#8BE9FD;font-style:italic;">Buffer</span></span>
<span class="line"><span style="color:#F8F8F2;">err </span><span style="color:#FF79C6;">:=</span><span style="color:#F8F8F2;"> tpl.</span><span style="color:#50FA7B;">Execute</span><span style="color:#F8F8F2;">(</span></span>
<span class="line"><span style="color:#F8F8F2;">    templatex.</span><span style="color:#50FA7B;">WithContent</span><span style="color:#F8F8F2;">(</span><span style="color:#E9F284;">&quot;</span><span style="color:#F1FA8C;">{{.Count}} items are made of {{.Material}}</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">),</span></span>
<span class="line"><span style="color:#F8F8F2;">    templatex.</span><span style="color:#50FA7B;">WithData</span><span style="color:#F8F8F2;">(</span><span style="color:#FF79C6;">map</span><span style="color:#F8F8F2;">[</span><span style="color:#8BE9FD;font-style:italic;">string</span><span style="color:#F8F8F2;">]</span><span style="color:#FF79C6;">interface</span><span style="color:#F8F8F2;">{}{</span></span>
<span class="line"><span style="color:#E9F284;">        &quot;</span><span style="color:#F1FA8C;">Count</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">:    </span><span style="color:#BD93F9;">3</span><span style="color:#F8F8F2;">,</span></span>
<span class="line"><span style="color:#E9F284;">        &quot;</span><span style="color:#F1FA8C;">Material</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">: </span><span style="color:#E9F284;">&quot;</span><span style="color:#F1FA8C;">wood</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">,</span></span>
<span class="line"><span style="color:#F8F8F2;">    }),</span></span>
<span class="line"><span style="color:#F8F8F2;">    templatex.</span><span style="color:#50FA7B;">WithIO</span><span style="color:#F8F8F2;">(</span><span style="color:#FF79C6;">&amp;</span><span style="color:#F8F8F2;">buf),</span></span>
<span class="line"><span style="color:#F8F8F2;">)</span></span>
<span class="line"><span style="color:#FF79C6;">if</span><span style="color:#F8F8F2;"> err </span><span style="color:#FF79C6;">!=</span><span style="color:#BD93F9;"> nil</span><span style="color:#F8F8F2;"> {</span></span>
<span class="line"><span style="color:#F8F8F2;">    log.</span><span style="color:#50FA7B;">Fatal</span><span style="color:#F8F8F2;">().</span><span style="color:#50FA7B;">Err</span><span style="color:#F8F8F2;">(err).</span><span style="color:#50FA7B;">Msg</span><span style="color:#F8F8F2;">(</span><span style="color:#E9F284;">&quot;</span><span style="color:#F1FA8C;">failed to execute template</span><span style="color:#E9F284;">&quot;</span><span style="color:#F8F8F2;">)</span></span>
<span class="line"><span style="color:#F8F8F2;">}</span></span>
<span class="line"></span>
<span class="line"><span style="color:#F8F8F2;">fmt.</span><span style="color:#50FA7B;">Println</span><span style="color:#F8F8F2;">(buf.</span><span style="color:#50FA7B;">String</span><span style="color:#F8F8F2;">())</span></span>
<span class="line"><span style="color:#6272A4;">// Output:</span></span>
<span class="line"><span style="color:#6272A4;">// 3 items are made of wood</span></span></code></pre></div>`,19)]))}const d=a(o,[["render",t]]);export{u as __pageData,d as default};