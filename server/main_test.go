package main_test

import (
	"context"

	pb "github.com/lucas.saraiva019/calculadora/proto/calculator"
	main "github.com/lucas.saraiva019/calculadora/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	calculator := main.Server{}
	Describe("Test the func of calculator", func() {
		LocalPattern1 := func(num1, num2, res float32, oper pb.OperatorType) {
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

		LocalPattern2 := func(num1, num2, res float32, oper pb.OperatorType) func() {
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
		Context("Operator Sum", func() {
			LocalPattern1(2, 2, 4, pb.OperatorType_SUM)
			LocalPattern1(2.7, 2.5, 5.2, pb.OperatorType_SUM)
			LocalPattern1(0, 0, 0, pb.OperatorType_SUM)
			LocalPattern1(-1, -2, -3, pb.OperatorType_SUM)
			LocalPattern1(-1.2, -2.3, -3.5, pb.OperatorType_SUM)
		})
		Context("Operator Sub", func() {
			It("Sub with two integer numbers", LocalPattern2(4, 2, 2, pb.OperatorType_SUBTRACTION))
			It("Sub with two floats numbers", LocalPattern2(4.5, 2.3, 2.2, pb.OperatorType_SUBTRACTION))
			It("Sub with two zero´s numbers", LocalPattern2(0, 0, 0, pb.OperatorType_SUBTRACTION))
			It("Sub with two integer negative numbers", LocalPattern2(-1, -2, 1, pb.OperatorType_SUBTRACTION))
			It("Sub with two float negative numbers", PatternComparerFloat2(-1.5, -2.3, 0.8, pb.OperatorType_SUBTRACTION))
		})
		Context("Operator Mult", func() {
			GlobalPattern1(4, 2, 8, pb.OperatorType_MULTIPLICATION)
			PatternComparerFloat1(4.5, 2.3, 10.35, pb.OperatorType_MULTIPLICATION)
			GlobalPattern1(0, 0, 0, pb.OperatorType_MULTIPLICATION)
			GlobalPattern1(-1, -2, 2, pb.OperatorType_MULTIPLICATION)
			PatternComparerFloat1(-1.5, -2.3, 3.45, pb.OperatorType_MULTIPLICATION)
		})
		Context("Operator Div", func() {
			It("Div with two integer numbers", GlobalPattern2(4, 2, 2, pb.OperatorType_DIVISION))
			It("Div with two floats numbers", PatternComparerFloat2(4.5, 2.3, 1.95652, pb.OperatorType_DIVISION))
			It("Div with two integer negative numbers", GlobalPattern2(-1, -2, 0.5, pb.OperatorType_DIVISION))
			It("Div with two float negative numbers", GlobalPattern2(-3.0, -1.5, 2, pb.OperatorType_DIVISION))
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
		})
	})
})
