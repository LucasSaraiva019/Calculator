package main_test

import (
	"context"
	"math"

	pb "example.com/calculadora/calculator"
	main "example.com/calculadora/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const tolerance = .00001

func ComparerFloat(x, y float64) bool {
	diff := math.Abs(x - y)
	mean := math.Abs(x+y) / 2.0
	return (diff / mean) < tolerance
}

var _ = Describe("Main", func() {
	calculator := main.Math{}
	Describe("Test the func of calculator", func() {
		Context("Operator Sum", func() {
			It("Sum with two integer numbers", func() {
				request := &pb.Request{
					NumberOne: 2,
					NumberTwo: 2,
					Operation: pb.OperatorType_SUM,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(4))
			})
			It("Sum with two float numbers", func() {
				request := &pb.Request{
					NumberOne: 2.7,
					NumberTwo: 2.5,
					Operation: pb.OperatorType_SUM,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(ComparerFloat(float64(response.Result), 5.2)).To(BeTrue())
			})
			It("Sum with two zero´s numbers", func() {
				request := &pb.Request{
					NumberOne: 0,
					NumberTwo: 0,
					Operation: pb.OperatorType_SUM,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(0))
			})
			It("Sum with two negative numbers", func() {
				request := &pb.Request{
					NumberOne: -1,
					NumberTwo: -2,
					Operation: pb.OperatorType_SUM,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(-3))
			})
		})
		Context("Operator Sub", func() {
			It("Sub with two integer numbers", func() {
				request := &pb.Request{
					NumberOne: 4,
					NumberTwo: 2,
					Operation: pb.OperatorType_SUBTRACTION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(2))
			})
			It("Sub with two floats numbers", func() {
				request := &pb.Request{
					NumberOne: 4.5,
					NumberTwo: 2.3,
					Operation: pb.OperatorType_SUBTRACTION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(ComparerFloat(float64(response.Result), 2.2)).To(BeTrue())
			})
		})
	})
})
