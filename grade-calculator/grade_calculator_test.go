package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeC(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("Assignment", 75, Assignment)
	gc.AddGrade("Exam", 75, Exam)
	gc.AddGrade("Essay", 75, Essay)
	if actual := gc.GetFinalGrade(); actual != "C" { t.Errorf("Expected C, got %s", actual)}

}

func TestGradeD(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("Assignment", 65, Assignment)
	gc.AddGrade("Exam", 65, Exam)
	gc.AddGrade("Essay", 65, Essay)
	if actual := gc.GetFinalGrade(); actual != "D" { t.Errorf("Expected D, got %s", actual)}
}

func TestGradeFailing(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("Assignment", 50, Assignment)
	gc.AddGrade("Exam", 50, Exam)
	gc.AddGrade("Essay", 49, Essay)
	if actual := gc.GetFinalGrade(); actual != "F" { t.Errorf("Expected F, got %s", actual)}
}

func TestEmpty(t *testing.T) {
	gc := NewGradeCalculator()
	if actual := gc.GetFinalGrade(); actual != "F" { t.Errorf("Expected F, got %s", actual)}
}

func TestGradeString(t *testing.T) {
	tests := []struct {
		gradeType GradeType
		expected      string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}
	for _, tt := range tests {
		if actual := tt.gradeType.String(); actual != tt.expected {
			t.Errorf("Expected %s, got %s", tt.expected, actual)
		}
	}
}

func TestMultGradesPerCategory(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("Assignment 1", 80, Assignment)
	gc.AddGrade("Assignment 2", 90, Assignment)
	gc.AddGrade("Exam 1", 85, Exam)
	gc.AddGrade("Exam 2", 95, Exam)
	gc.AddGrade("Essay 1", 75, Essay)
	gc.AddGrade("Essay 2", 85, Essay)
	if actual := gc.GetFinalGrade(); actual != "B" { t.Errorf("Expected B, got %s", actual)}
}