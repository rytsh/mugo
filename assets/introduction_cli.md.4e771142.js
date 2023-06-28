import{_ as s,o as a,c as n,O as l}from"./chunks/framework.cdae9d71.js";const f=JSON.parse('{"title":"CLI","description":"","frontmatter":{},"headers":[],"relativePath":"introduction/cli.md","filePath":"introduction/cli.md"}'),e={name:"introduction/cli.md"},t=l(`<h1 id="cli" tabindex="-1">CLI <a class="header-anchor" href="#cli" aria-label="Permalink to &quot;CLI&quot;">​</a></h1><div class="language-"><button title="Copy Code" class="copy"></button><span class="lang"></span><pre class="shiki material-theme-palenight"><code><span class="line"><span style="color:#A6ACCD;">Usage:</span></span>
<span class="line"><span style="color:#A6ACCD;">  mugo &lt;template&gt; [flags]</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">Examples:</span></span>
<span class="line"><span style="color:#A6ACCD;">mugo -d @data.yaml template.tpl</span></span>
<span class="line"><span style="color:#A6ACCD;">mugo -d &#39;{&quot;Name&quot;: &quot;mugo&quot;}&#39; -o output.txt template.tpl</span></span>
<span class="line"><span style="color:#A6ACCD;">mugo -d &#39;{&quot;Name&quot;: &quot;mugo&quot;}&#39; -o output.txt - &lt; template.tpl</span></span>
<span class="line"><span style="color:#A6ACCD;">mugo -d &#39;{&quot;Name&quot;: &quot;mugo&quot;}&#39; - &lt;&lt;&lt; &quot;{{.Name}}&quot;</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span>
<span class="line"><span style="color:#A6ACCD;">Flags:</span></span>
<span class="line"><span style="color:#A6ACCD;">  -d, --data stringArray            input data as json/yaml or file path with @ prefix could be &#39;.yaml&#39;,&#39;.yml&#39;,&#39;.json&#39;,&#39;.toml&#39; extension</span></span>
<span class="line"><span style="color:#A6ACCD;">  -r, --data-raw string             input data as raw or file path with @ prefix could be file with any extension</span></span>
<span class="line"><span style="color:#A6ACCD;">      --delims string               comma or space separated list of delimiters to alternate the default &quot;{{ }}&quot;</span></span>
<span class="line"><span style="color:#A6ACCD;">      --disable-func stringArray    disabled functions for run template</span></span>
<span class="line"><span style="color:#A6ACCD;">      --disable-group stringArray   disabled groups for run template</span></span>
<span class="line"><span style="color:#A6ACCD;">      --enable-func stringArray     specific functions for run template</span></span>
<span class="line"><span style="color:#A6ACCD;">      --enable-group stringArray    specific function groups for run template</span></span>
<span class="line"><span style="color:#A6ACCD;">  -h, --help                        help for mugo</span></span>
<span class="line"><span style="color:#A6ACCD;">      --html                        use html/template instead</span></span>
<span class="line"><span style="color:#A6ACCD;">  -k, --insecure                    skip verify ssl certificate</span></span>
<span class="line"><span style="color:#A6ACCD;">  -l, --list                        function list</span></span>
<span class="line"><span style="color:#A6ACCD;">      --log-level string            log level (debug, info, warn, error, fatal, panic), default is info (default &quot;info&quot;)</span></span>
<span class="line"><span style="color:#A6ACCD;">      --no-at                       disable @ prefix for file path</span></span>
<span class="line"><span style="color:#A6ACCD;">      --no-retry                    disable retry</span></span>
<span class="line"><span style="color:#A6ACCD;">  -o, --output string               output file, default is stdout</span></span>
<span class="line"><span style="color:#A6ACCD;">  -p, --parse stringArray           parse file pattern for define templates &#39;testdata/**/*.tpl&#39;</span></span>
<span class="line"><span style="color:#A6ACCD;">      --perm-file string            create file permission, default is 0644</span></span>
<span class="line"><span style="color:#A6ACCD;">      --perm-folder string          create folder permission, default is 0755</span></span>
<span class="line"><span style="color:#A6ACCD;">  -s, --silience                    silience log</span></span>
<span class="line"><span style="color:#A6ACCD;">  -t, --trust                       trust to execute dangerous functions</span></span>
<span class="line"><span style="color:#A6ACCD;">  -v, --version                     version for mugo</span></span>
<span class="line"><span style="color:#A6ACCD;">  -w, --work-dir string             work directory for run template</span></span></code></pre></div>`,2),p=[t];function o(r,i,c,u,A,C){return a(),n("div",null,p)}const m=s(e,[["render",o]]);export{f as __pageData,m as default};
