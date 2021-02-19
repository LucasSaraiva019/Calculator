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

var _ = Describe("Main", func() {
	calculator := main.Server{}
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
			It("Sum with two integer negative numbers", func() {
				request := &pb.Request{
					NumberOne: -1,
					NumberTwo: -2,
					Operation: pb.OperatorType_SUM,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(-3))
			})
			It("Sum with two float negative numbers", func() {
				request := &pb.Request{
					NumberOne: -1.2,
					NumberTwo: -2.3,
					Operation: pb.OperatorType_SUM,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(ComparerFloat((float64(response.Result)), -3.5)).To(BeTrue())
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
			It("Sub with two zero´s numbers", func() {
				request := &pb.Request{
					NumberOne: 0,
					NumberTwo: 0,
					Operation: pb.OperatorType_SUBTRACTION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(0))
			})
			It("Sub with two integer negative numbers", func() {
				request := &pb.Request{
					NumberOne: -1,
					NumberTwo: -2,
					Operation: pb.OperatorType_SUBTRACTION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(1))
			})
			It("Sub with two float negative numbers", func() {
				request := &pb.Request{
					NumberOne: -1.5,
					NumberTwo: -2.3,
					Operation: pb.OperatorType_SUBTRACTION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(ComparerFloat(float64(response.Result), 0.8)).To(BeTrue())
			})
		})
		Context("Operator Mult", func() {
			It("Mult with two integer numbers", func() {
				request := &pb.Request{
					NumberOne: 4,
					NumberTwo: 2,
					Operation: pb.OperatorType_MULTIPLICATION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(8))
			})
			It("Sub with two floats numbers", func() {
				request := &pb.Request{
					NumberOne: 4.5,
					NumberTwo: 2.3,
					Operation: pb.OperatorType_MULTIPLICATION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(ComparerFloat(float64(response.Result), 10.35)).To(BeTrue())
			})
			It("Sub with two zero´s numbers", func() {
				request := &pb.Request{
					NumberOne: 0,
					NumberTwo: 0,
					Operation: pb.OperatorType_MULTIPLICATION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(0))
			})
			It("Sub with two integer negative numbers", func() {
				request := &pb.Request{
					NumberOne: -1,
					NumberTwo: -2,
					Operation: pb.OperatorType_MULTIPLICATION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(2))
			})
			It("Sub with two float negative numbers", func() {
				request := &pb.Request{
					NumberOne: -1.5,
					NumberTwo: -2.3,
					Operation: pb.OperatorType_MULTIPLICATION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(ComparerFloat(float64(response.Result), 3.45)).To(BeTrue())
			})
		})
		Context("Operator Div", func() {
			It("Div with two integer numbers", func() {
				request := &pb.Request{
					NumberOne: 4,
					NumberTwo: 2,
					Operation: pb.OperatorType_DIVISION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(2))
			})
			It("Div with two floats numbers", func() {
				request := &pb.Request{
					NumberOne: 4.5,
					NumberTwo: 2.3,
					Operation: pb.OperatorType_DIVISION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(ComparerFloat(float64(response.Result), 1.95652)).To(BeTrue())
			})
			It("Div with second number is zero numbers", func() {
				request := &pb.Request{
					NumberOne: 10,
					NumberTwo: 0,
					Operation: pb.OperatorType_DIVISION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("Não é possivel dividir por Zero"))
				Expect(response).To(BeNil())
			})
			It("Div with two integer negative numbers", func() {
				request := &pb.Request{
					NumberOne: -1,
					NumberTwo: -2,
					Operation: pb.OperatorType_DIVISION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(response.Result).To(BeEquivalentTo(0.5))
			})
			It("Div with two float negative numbers", func() {
				request := &pb.Request{
					NumberOne: -3.0,
					NumberTwo: -1.5,
					Operation: pb.OperatorType_DIVISION,
				}
				response, err := calculator.Calculate(context.Background(), request)
				Expect(err).To(BeNil())
				Expect(ComparerFloat(float64(response.Result), 2)).To(BeTrue())
			})
		})
	})
})
