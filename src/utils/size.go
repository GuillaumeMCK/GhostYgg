package utils

type Size struct {
	Width  int
	Height int
}

func NewSize(width, height int) Size {
	return Size{
		Width:  width,
		Height: height,
	}
}

func (s *Size) Resize(width, height int) *Size {
	s.Width = width
	s.Height = height
	return s
}

func (s *Size) ResizeWidth(width int) *Size {
	s.Width = width
	return s
}

func (s *Size) ResizeHeight(height int) *Size {
	s.Height = height
	return s
}
