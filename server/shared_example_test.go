package main_test

import (
	"context"
	"math"

	pb "github.com/lucas.saraiva019/calculadora/proto/calculator"
	main "github.com/lucas.saraiva019/calculadora/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const tolerance = .00001

func ComparerFloat(x, y float64) bool {
	diff := math.Abs(x - y)
	mean := math.Abs(x+y) / 2.0
	return (diff / mean) < tolerance
}

func GlobalPattern1(num1, num2, res float32, oper pb.OperatorType) {
	calculator := main.Server{}
	It("Returns the sum,sub,div and mult", func() {
		request := &pb.Request{
			NumberOne: num1,
			NumberTwo: num2,
			Operation: oper,
		}
		response, err := calculator.Calculate(context.Background(), request)
		Expect(err).To(BeNil())
		Expect(response.Result).To(BeEquivalentTo(res))

	})
}

func PatternComparerFloat1(num1, num2, res float32, oper pb.OperatorType) {
	calculator := main.Server{}
	It("Returns the sum,sub,div and mult with comparer Float", func() {
		request := &pb.Request{
			NumberOne: num1,
			NumberTwo: num2,
			Operation: oper,
		}
		response, err := calculator.Calculate(context.Background(), request)
		Expect(err).To(BeNil())
		Expect(ComparerFloat(float64(response.Result), float64(res))).To(BeTrue())

	})
}

func GlobalPattern2(num1, num2, res float32, oper pb.OperatorType) func() {
	calculator := main.Server{}
	return func() {
		request := &pb.Request{
			NumberOne: num1,
			NumberTwo: num2,
			Operation: oper,
		}
		response, err := calculator.Calculate(context.Background(), request)
		Expect(err).To(BeNil())
		Expect(response.Result).To(BeEquivalentTo(res))
	}
}

func PatternComparerFloat2(num1, num2, res float32, oper pb.OperatorType) func() {
	calculator := main.Server{}
	return func() {
		request := &pb.Request{
			NumberOne: num1,
			NumberTwo: num2,
			Operation: oper,
		}
		response, err := calculator.Calculate(context.Background(), request)
		Expect(err).To(BeNil())
		Expect(ComparerFloat(float64(response.Result), float64(res))).To(BeTrue())

	}
}
