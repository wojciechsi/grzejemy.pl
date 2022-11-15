package controllers

import "github.com/wojciechsi/grzejemy.pl/models"

func GetOfferDetails(offer models.Offer) string {
	return offer.GetName()
}
