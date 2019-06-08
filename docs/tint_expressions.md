### Tint Color Expressions (TCE)

In `v0.0.1` tint introduced a new way to style your output. This solution was better than `t.Palette()` as it was
declared inline and did not have to create new Mixin instances.

The below representation of tint covers the expression structure:

| **`Feature`** | **`Prefix`** | **`Suffix`** |
| ------------ | --------------| ------------|
| Foreground | <code>{color}&#124;</code> | <code>&#124;!</code>  |
| Foreground-Bold | <code>*{color}&#124;</code> | <code>&#124;!</code>  |
| Background | <code>+{color}&#124;</code> | <code>&#124;+</code>  |

