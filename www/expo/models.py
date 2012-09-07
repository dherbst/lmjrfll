from django.db import models


class Expo(models.Model):
    '''
    The expo is the date, times, and location where the teams will present their projects.
    '''
    name = models.CharField(max_length=100)
    date = models.DateField()
    start_time = models.TimeField()
    end_time = models.TimeField()

    def __unicode__(self):
        return self.name
    


class Team(models.Model):
    '''
    A JRFLL team presenting at the expo.
    '''
    name = models.CharField(max_length=100)
    jrfll_team_number = models.IntegerField()
    is_looking_for_members = models.BooleanField(default=False)
    
    def __unicode__(self):
        return self.name
    
    
    
