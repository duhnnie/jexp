package expression

import "fmt"

func getPath(t, currentPath, childPath string) string {
	if t == expTypeVar {
		return fmt.Sprintf("%s:%s", t, currentPath)
	}

	return fmt.Sprintf("%s.%s%s", t, currentPath, childPath)
}
