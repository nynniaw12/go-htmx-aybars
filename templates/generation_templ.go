// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Gen() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body><div class=\"space-y-4 text-white mx-[20px] mt-[20px] flex flex-col\"><h1 class=\"text-3xl font-bold mb-4\">Heading 1</h1><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, nisl nec tincidunt lacinia, nisl nisl aliquam nisl, eget aliquam nisl nisl eget nisl.</p><h2 class=\"text-xl font-bold mb-2\">Heading 2</h2><p>Sed auctor neque sed ipsum bibendum feugiat. Fusce commodo ultricies nisi, eu molestie velit sagittis eget.</p><img class=\"max-w-full h-auto max-h-80 mx-auto\" src=\"https://via.placeholder.com/800x600\" alt=\"Sample Image\"><table class=\"text-white table-auto divide-y divide-gray-700 divide-x divide-gray-700 border-2 border-gray-700\"><tr><th>Column 1</th><th>Column 2</th><th>Column 3</th></tr><tr><td>Value 1</td><td>Value 2</td><td>Value 3</td></tr><tr><td>Value 4</td><td>Value 5</td><td>Value 6</td></tr></table></div></body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
