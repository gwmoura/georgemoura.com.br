# Import the Flask Framework
from flask import Flask, url_for, render_template
import datetime
app = Flask(__name__)
# Note: We don't need to call run() since our application is embedded within
# the App Engine WSGI application server.

@app.route('/')
@app.route('/curriculo')
@app.route('/curriculo/')
def hello():
  today = datetime.date.today()
  birthday = datetime.date(day=14, month=5, year=1990)
  diff = today - birthday
  years_old = diff.days/365
  return render_template('cv.html', years_old = years_old)


@app.errorhandler(404)
def page_not_found(e):
  """Return a custom 404 error."""
  return 'Sorry, Nothing at this URL.', 404


@app.errorhandler(500)
def application_error(e):
  """Return a custom 500 error."""
  return 'Sorry, unexpected error: {}'.format(e), 500